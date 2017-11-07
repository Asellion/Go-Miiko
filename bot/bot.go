package bot

import (
	"fmt"

	"../config"
	"github.com/bwmarrin/discordgo"
)

// BotID : Numerical ID of the bot
var BotID string
var goBot *discordgo.Session

// Start : Starts the bot.
func Start() {

	// Go online!
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get Bot ID
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID

	// Hey, listen!
	goBot.AddHandler(messageHandler)

	// Crash on error
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// It's alive!
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Myself?
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
