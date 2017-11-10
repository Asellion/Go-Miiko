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
	goBot.AddHandler(reactHandler)
	goBot.AddHandler(leaveHandler)

	// Crash on error
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// It's alive!
	fmt.Println("Miiko is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Myself?
	if m.Author.ID == BotID {
		return
	}

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err.Error())
	}

	// DM?
	if channel.Type == discordgo.ChannelTypeDM {
		if config.BotMasterChannelID == "" {

			// No BotMaster
			fmt.Println(m.Author.Username + " : " + m.Content)

		} else if m.ChannelID == config.BotMasterChannelID {
			// Talking to BotMaster
		} else {

			// Typing!
			err = s.ChannelTyping(config.BotMasterChannelID)
			if err != nil {
				fmt.Println(err.Error())
			}

			// Foward the message to BotMaster!
			_, err := s.ChannelMessageSend(config.BotMasterChannelID, "<@"+m.Author.ID+"> : "+m.Content)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		return
	}

	// Ask for guard
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		askForGuard(s, m)
		return
	}

	// Place in a guard
	placeInAGuard(s, m)
}

func reactHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Pin popular message
	pin(s, m)
}

func leaveHandler(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Invite people who leave
	waitComeBack(s, m)
}
