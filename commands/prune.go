package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var pruning = make(map[string]bool)

// Prune a server from inactive people with a role
func Prune(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get a channel structure.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Get guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get a guild structure.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	if m.Author.ID != guild.OwnerID {
		return
	}

	// Check if already pruning
	alreadyPruning, exists := pruning[guild.ID]
	if exists {
		if alreadyPruning {

			// Stop!
			s.ChannelTyping(channel.ID)
			_, err = s.ChannelMessageSend(channel.ID, "Désolée <@"+m.Author.ID+">! Je purge déjà la guilde "+guild.Name+".")
			if err != nil {
				fmt.Println("Couldn't send a message.")
				fmt.Println("Guild : " + guild.Name)
				fmt.Println("Channel : " + channel.Name)
				fmt.Println("Author : " + m.Author.Username)
				fmt.Println("Message : " + m.Content)
				fmt.Println(err.Error())
			}
			return
		}
	}

	// Pruning for this server!
	pruning[guild.ID] = true
	start := time.Now()

	// Announce
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "<@"+m.Author.ID+"> Début de la purification de "+guild.Name+"! Ça peut prendre quelques minutes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
	}

	// List of members and their role
	var MembersMap = make(map[string][]string)

	// For all members
	for _, gMember := range guild.Members {

		// Save their roles
		MembersMap[gMember.User.ID] = gMember.Roles

		// Remove their roles
		err := s.GuildMemberEdit(guild.ID, gMember.User.ID, []string{})
		if err != nil {
			fmt.Println("Couldn't remove roles.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Member : " + gMember.User.Username)
			fmt.Println(err.Error())
		}
	}

	// Prune
	count, err := s.GuildPrune(guild.ID, 30)
	sCount := strconv.Itoa(int(count))

	// For all members
	for _, gMember := range guild.Members {

		// Give back their roles
		err = s.GuildMemberEdit(guild.ID, gMember.User.ID, MembersMap[gMember.User.ID])
		if err != nil {
			fmt.Println("Couldn't give back roles.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Member : " + gMember.User.Username)
			fmt.Println(err.Error())
		}
	}

	// Get task duration
	elapsed := time.Since(start)
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) - minutes*60

	// Over!
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, getPruneMessage(sCount))
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println(err.Error())
	}

	// Send task duration
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "Opération terminée en "+strconv.Itoa(minutes)+" minutes et "+strconv.Itoa(seconds)+" secondes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println(err.Error())
	}

	// Stop pruning for this server.
	pruning[guild.ID] = false
}

func getPruneMessage(sCount string) string {

	// Prune Messages
	var pruneList []string
	pruneList = append(pruneList, sCount+" inactifs ont été kickés.")
	pruneList = append(pruneList, "Le serveur a été purifié de "+sCount+" inactifs.")

	// Seed
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return pruneList[rand.Intn(len(pruneList))]
}
