package bot

import (
	"fmt"
	"strings"

	"github.com/NatoBoram/Go-Miiko/config"
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
	goBot.AddHandler(joinHandler)

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
	if m.Author.ID == BotID || !m.Author.Bot {
		return
	}

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get guild structure
	guild, err := s.State.Guild(channel.GuildID)
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
		config.UpdateWelcomeChannel(s, m)
		return
	}

	// Place in a guard
	placeInAGuard(s, m)

	// Mentionned someone?
	if len(m.Mentions) > 0 {
		for x := 0; x < len(m.Mentions); x++ {

			// Mentionned me?
			if m.Mentions[x].ID == BotID {

				// Command Set Welcome Channel
				if (strings.Contains(m.Content, "set") && strings.Contains(m.Content, "welcome") && strings.Contains(m.Content, "channel") && m.Author.ID == guild.OwnerID) && !strings.Contains(m.Content, "\\") {
					config.UpdateWelcomeChannel(s, m)
					_, err := s.ChannelMessageSend(m.ChannelID, "D'accord! Ce channel est maintenant le channel de bienvenue.")
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
		}
	}

	// Popcorn?
	popcorn(s, m)
}

func reactHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Pin popular message
	pin(s, m)
}

func leaveHandler(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Invite people who leave
	waitComeBack(s, m)
}

func joinHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// Myself?
	if m.User.ID != BotID {

		// Ask for guard
		askForGuard(s, m)
	}
}
