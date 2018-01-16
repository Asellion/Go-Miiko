package bot

import (
	"fmt"
	"strings"

	"github.com/NatoBoram/Go-Miiko/commands"
	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
)

// BotID : Numerical ID of the bot
var BotID string
var goBot *discordgo.Session

// Start : Starts the bot.
func Start() {

	// Go online!
	goBot, err := discordgo.New("Bot " + config.Database.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get Bot ID
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Couldn't get the BotID.")
		fmt.Println(err.Error())
		return
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
	fmt.Println("Hi, Master. I am Miiko, and everything's all right!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Myself?
	if m.Author.ID == BotID {
		return
	}

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

	// DM?
	if channel.Type == discordgo.ChannelTypeDM {

		// Popcorn?
		commands.Popcorn(s, m)

		if config.Database.MasterID == "" {

			// No BotMaster
			fmt.Println(m.Author.Username + " : " + m.Content)

		} else if m.Author.ID == config.Database.MasterID {

			// Talking to Master

		} else {

			// Get Master's User
			user, err := s.User(config.Database.MasterID)
			if err != nil {
				fmt.Println("Couldn't get Master's User!")
				fmt.Println(err.Error())
				return
			}

			// Create channel with Master
			masterChannel, err := s.UserChannelCreate(config.Database.MasterID)
			if err != nil {
				fmt.Println("Couldn't create a private channel with " + user.Username + ".")
				fmt.Println(err.Error())
				return
			}

			// Foward the message to BotMaster!
			s.ChannelTyping(masterChannel.ID)
			_, err = s.ChannelMessageSend(masterChannel.ID, "<@"+m.Author.ID+"> : "+m.Content)
			if err != nil {
				fmt.Println("Couldn't foward a message to " + user.Username + ".")
				fmt.Println("Author : " + m.Author.Username)
				fmt.Println("Message : " + m.Content)
				fmt.Println(err.Error())
				return
			}
		}
		return
	}

	// Update welcome channel
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		config.UpdateWelcomeChannel(s, m)
		return
	}

	// Bot?
	if m.Author.Bot {
		return
	}

	// Place in a guard
	placeInAGuard(s, m)

	// Mentionned someone?
	if len(m.Mentions) > 0 {
		for x := 0; x < len(m.Mentions); x++ {

			// Mentionned me?
			if m.Mentions[x].ID == BotID {

				// Split
				command := strings.Split(m.Content, " ")

				// Commands with 2 words
				if len(command) == 2 {
					if command[1] == "prune" {
						commands.Prune(s, m)
					}
				}

				// Commands with 3 words
				if len(command) == 3 {
					if command[1] == "get" {
						if command[2] == "points" {
							commands.GetPoints(s, m)
						}
					}
				}

				// Commands with 4 words
				if len(command) == 4 {
					if command[1] == "set" {
						if command[2] == "welcome" {
							if command[3] == "channel" && m.Author.ID == guild.OwnerID {
								commands.SetWelcomeChannel(s, m)
							}
						}
					}
					if command[1] == "get" {
						if command[2] == "welcome" {
							if command[3] == "channel" {
								commands.GetWelcomeChannel(s, m)
							}
						}
					}
				}
			}
		}
	}

	// Reactions
	commands.Popcorn(s, m)
	commands.Nani(s, m)
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
