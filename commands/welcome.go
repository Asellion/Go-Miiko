package commands

import (
	"fmt"

	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
)

// SetWelcomeChannel sets the welcome channel
func SetWelcomeChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a said message.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	s.ChannelTyping(channel.ID)
	config.UpdateWelcomeChannel(s, m)
	_, err = s.ChannelMessageSend(channel.ID, "D'accord! Ce salon est maintenant le salon de bienvenue.")
	if err != nil {
		fmt.Println("Couldn't send a message in " + channel.Name + ".")
		fmt.Println(err.Error())
	}
}

// GetWelcomeChannel gets the welcome channel
func GetWelcomeChannel(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a said message.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Get guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild structure of a said message.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	s.ChannelTyping(channel.ID)
	_, err = s.ChannelMessageSend(channel.ID, "Le salon de bienvenue est <#"+config.GetWelcomeChannelByGuildID(guild.ID)+">.")
	if err != nil {
		fmt.Println("Couldn't send a message in " + channel.Name + ".")
		fmt.Println(err.Error())
	}
}
