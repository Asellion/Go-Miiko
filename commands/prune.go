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
func Prune(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {

	if m.Author.ID != g.OwnerID {
		return
	}

	// Check if already pruning
	alreadyPruning, exists := pruning[g.ID]
	if exists {
		if alreadyPruning {

			// Stop!
			s.ChannelTyping(c.ID)
			_, err := s.ChannelMessageSend(c.ID, "Désolée <@"+m.Author.ID+">! Je purge déjà la guilde "+g.Name+".")
			if err != nil {
				fmt.Println("Couldn't send a message.")
				fmt.Println("Guild : " + g.Name)
				fmt.Println("Channel : " + c.Name)
				fmt.Println("Author : " + m.Author.Username)
				fmt.Println("Message : " + m.Content)
				fmt.Println(err.Error())
			}
			return
		}
	}

	// Announce
	s.ChannelTyping(c.ID)
	_, err := s.ChannelMessageSend(c.ID, "<@"+m.Author.ID+"> Début de la purification de "+g.Name+"! Ça peut prendre quelques minutes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Pruning for this server!
	pruning[g.ID] = true
	start := time.Now()

	// List of members and their role
	var MembersMap = make(map[string][]string)

	// For all members
	for _, gMember := range g.Members {

		// Save their roles
		MembersMap[gMember.User.ID] = gMember.Roles

		// Remove their roles
		err := s.GuildMemberEdit(g.ID, gMember.User.ID, []string{})
		if err != nil {
			fmt.Println("Couldn't remove roles.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + gMember.User.Username)
			fmt.Println(err.Error())
		}
	}

	// Prune
	count, err := s.GuildPrune(g.ID, 30)
	sCount := strconv.Itoa(int(count))

	// For all members
	for _, gMember := range g.Members {

		// Give back their roles
		err = s.GuildMemberEdit(g.ID, gMember.User.ID, MembersMap[gMember.User.ID])
		if err != nil {
			fmt.Println("Couldn't give back roles.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + gMember.User.Username)
			fmt.Println(err.Error())
		}
	}

	// Get task duration
	elapsed := time.Since(start)
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) - minutes*60

	// Over!
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, getPruneMessage(sCount))
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
	}

	// Send task duration
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Opération terminée en "+strconv.Itoa(minutes)+" minutes et "+strconv.Itoa(seconds)+" secondes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
	}

	// Stop pruning for this server.
	pruning[g.ID] = false
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
