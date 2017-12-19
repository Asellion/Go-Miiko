package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var counting = false

// GetPoints counts the points of a server
func GetPoints(s *discordgo.Session, m *discordgo.MessageCreate) {

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

	// Anti-Spam
	if counting {
		s.ChannelTyping(channel.ID)
		_, err := s.ChannelMessageSend(channel.ID, "Désolée <@"+m.Author.ID+">! Je suis déjà en train de compter des points. Réessaie dans quelques minutes!")
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

	// Announce
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "<@"+m.Author.ID+"> Je compte les points de "+guild.Name+"! Ça peut prendre quelques minutes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Variables
	counting = true
	points := make(map[string]int)
	start := time.Now()

	// Create feedback message
	s.ChannelTyping(channel.ID)
	feedback, err := s.ChannelMessageSend(channel.ID, "Je suis à 0%.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// For every channels
	for gIndex, gChannel := range guild.Channels {

		// Edit feedback message
		s.ChannelTyping(channel.ID)
		progress := 100 * gIndex / len(guild.Channels)
		_, err := s.ChannelMessageEdit(channel.ID, feedback.ID, "Je suis à "+strconv.Itoa(progress)+"%.")
		if err != nil {
			fmt.Println("Couldn't edit a message.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Channel : " + channel.Name)
			fmt.Println("Message : " + feedback.Content)
			fmt.Println(err.Error())
		}

		// Pinned messages are obviously only in text channels.
		if gChannel.Type != discordgo.ChannelTypeGuildText {
			continue
		}

		// Get every pinned messages
		cMessages, err := s.ChannelMessagesPinned(gChannel.ID)
		if err != nil {
			fmt.Println("Couldn't get pinned messages.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Channel : " + gChannel.Name)
			fmt.Println(err.Error())
			continue
		}

		// For every messages
		for _, cMessage := range cMessages {

			// Get the author
			member, err := s.GuildMember(guild.ID, cMessage.Author.ID)
			if err != nil {
				fmt.Println("Couldn't get a member.")
				fmt.Println("Guild : " + guild.Name)
				fmt.Println("Channel : " + gChannel.Name)
				fmt.Println("Author : " + cMessage.Author.Username)
				fmt.Println("Message : " + cMessage.Content)
				fmt.Println(err.Error())
				continue
			}

			// If the author has only one single role
			if len(member.Roles) == 1 {
				points[member.Roles[0]]++
			}
		}
	}

	// Delete feedback message
	err = s.ChannelMessageDelete(channel.ID, feedback.ID)
	if err != nil {
		fmt.Println("Couldn't delete a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + feedback.Author.Username)
		fmt.Println("Message : " + feedback.Content)
		fmt.Println(err.Error())
	}

	// Show points
	for key, value := range points {
		_, err := s.ChannelMessageSend(channel.ID, "<@&"+key+"> : "+strconv.Itoa(value))
		if err != nil {
			fmt.Println("Couldn't send a message.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Channel : " + channel.Name)
			fmt.Println(err.Error())
			continue
		}
	}

	// Get task duration
	elapsed := time.Since(start)
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) - minutes*60

	// Send task duration
	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "Opération terminée en "+strconv.Itoa(minutes)+" minutes et "+strconv.Itoa(seconds)+" secondes.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println(err.Error())
	}

	// Unlock this feature
	counting = false
}
