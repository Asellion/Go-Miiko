package bot

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
)

// BotID : Numerical ID of the bot
var BotID string
var goBot *discordgo.Session
var counting = false

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
		popcorn(s, m)

		if config.BotMasterChannelID == "" {

			// No BotMaster
			fmt.Println(m.Author.Username + " : " + m.Content)

		} else if m.ChannelID == config.BotMasterChannelID {

			// Talking to Master

		} else {

			// Foward the message to BotMaster!
			s.ChannelTyping(config.BotMasterChannelID)
			_, err := s.ChannelMessageSend(config.BotMasterChannelID, "<@"+m.Author.ID+"> : "+m.Content)
			if err != nil {
				fmt.Println("Couldn't foward a message to Master.")
				fmt.Println("Author : " + m.Author.Username)
				fmt.Println("Message : " + m.Content)
				fmt.Println(err.Error())
			}
		}
		return
	}

	// Update welcome channel
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		config.UpdateWelcomeChannel(s, m)
		return
	}

	// Popcorn?
	popcorn(s, m)

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

				// Commands with 3 words
				if len(command) == 3 {
					if command[1] == "get" {
						if command[2] == "points" {

							// Anti-Spam
							if counting {
								s.ChannelTyping(channel.ID)
								_, err := s.ChannelMessageSend(channel.ID, "Désolée <@"+m.Author.ID+">! Je suis déjà en train de compter des points. Réessaie dans quelques minutes!")
								if err != nil {
									fmt.Println("Couldn't send a message in " + channel.Name + ".")
									fmt.Println(err.Error())
								}
								return
							}

							// Announce
							s.ChannelTyping(channel.ID)
							_, err := s.ChannelMessageSend(channel.ID, "<@"+m.Author.ID+"> Je compte les points de "+guild.Name+"! Ça peut prendre quelques minutes.")
							if err != nil {
								fmt.Println("Couldn't send a message in " + channel.Name + ".")
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
								fmt.Println("Couldn't send a message in " + channel.Name + ".")
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
									fmt.Println("Couldn't edit a message in " + channel.Name + ".")
									fmt.Println(err.Error())
									return
								}

								// Pinned messages are obviously only in text channels.
								if gChannel.Type != discordgo.ChannelTypeGuildText {
									continue
								}

								// Get every pinned messages
								messages, err := s.ChannelMessagesPinned(gChannel.ID)
								if err != nil {
									fmt.Println("Couldn't get pinned messages of ", gChannel.Name, ".")
									fmt.Println(err.Error())
									continue
								}

								// For every messages
								for _, message := range messages {

									// Get the author
									member, err := s.GuildMember(guild.ID, message.Author.ID)
									if err != nil {
										fmt.Println("Couldn't get the member ", message.Author.Username, " in guild ", guild.Name, ".")
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
								fmt.Println("Couldn't delete a message in " + channel.Name + ".")
								fmt.Println(err.Error())
							}

							// Show points
							for key, value := range points {

								_, err := s.ChannelMessageSend(channel.ID, "<@&"+key+"> : "+strconv.Itoa(value))
								if err != nil {
									fmt.Println("Couldn't send a message in " + channel.Name + ".")
									fmt.Println(err.Error())
									continue
								}
							}

							// Task over.

							elapsed := time.Since(start)
							minutes := int(elapsed.Minutes())
							seconds := int(elapsed.Seconds()) - minutes*60

							s.ChannelTyping(channel.ID)
							_, err = s.ChannelMessageSend(channel.ID, "Opération terminée en "+strconv.Itoa(minutes)+" minutes et "+strconv.Itoa(seconds)+" secondes.")
							if err != nil {
								fmt.Println("Couldn't send a message in " + channel.Name + ".")
								fmt.Println(err.Error())
							}

							// Unlock this feature
							counting = false
						}
					}
				}

				// Commands with 4 words
				if len(command) == 4 {
					if command[1] == "set" {
						if command[2] == "welcome" {
							if command[3] == "channel" && m.Author.ID == guild.OwnerID {

								// Set welcome channel
								s.ChannelTyping(channel.ID)
								config.UpdateWelcomeChannel(s, m)
								_, err := s.ChannelMessageSend(channel.ID, "D'accord! Ce salon est maintenant le salon de bienvenue.")
								if err != nil {
									fmt.Println("Couldn't send a message in " + channel.Name + ".")
									fmt.Println(err.Error())
								}
							}
						}
					}
					if command[1] == "get" {
						if command[2] == "welcome" {
							if command[3] == "channel" {

								// Get welcome channel
								s.ChannelTyping(channel.ID)
								_, err := s.ChannelMessageSend(channel.ID, "Le salon de bienvenue est <#"+config.GetWelcomeChannelByGuildID(guild.ID)+">.")
								if err != nil {
									fmt.Println("Couldn't send a message in " + channel.Name + ".")
									fmt.Println(err.Error())
								}
							}
						}
					}
				}
			}
		}
	}
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
