package bot

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/NatoBoram/Go-Miiko/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	// DB : Connection to the database.
	DB *sql.DB

	// Me : The bot itself.
	Me *discordgo.User

	// Master : UserID of the bot's master.
	Master *discordgo.User
)

// Start : Starts the bot.
func Start(db *sql.DB, session *discordgo.Session, master string) error {

	// Database
	DB = db

	// Myself
	user, err := session.User("@me")
	if err != nil {
		fmt.Println("Couldn't get myself.")
		return err
	}
	Me = user

	// Master
	user, err = session.User(master)
	if err != nil {
		fmt.Println("Couldn't recognize my master.")
		return err
	}
	Master = user

	// Hey, listen!
	session.AddHandler(messageHandler)
	session.AddHandler(reactHandler)
	session.AddHandler(leaveHandler)
	session.AddHandler(joinHandler)

	// Refresh
	go refresh(db, session)

	// It's alive!
	fmt.Println("Hi, master " + Master.Username + ". I am " + Me.Username + ", and everything's all right!")

	// Everything is fine!
	return nil
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Myself? Super User?
	if m.Author.ID == Me.ID || m.Author.Discriminator == "0000" {
		return
	}

	// Functions 2.0
	done := false

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get a channel structure.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Forward to Master.
	done = forward(s, channel, m.Message)
	if done {
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

	// Get guild member
	member, err := s.GuildMember(channel.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("Couldn't get a member structure.")
		fmt.Println("Guild : " + guild.Name)
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
		return
	}

	// Update welcome channel
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		commands.SetWelcomeChannel(DB, s, guild, channel)
		return
	}

	// Guard
	done = commands.PlaceInAGuard(s, guild, channel, member, m.Message)
	if done {
		return
	}

	// Nani?!
	done = commands.Nani(s, m.Message)
	if done {
		return
	}

	// Popcorn!
	done = commands.Popcorn(s, channel, m.Message)
	if done {
		return
	}

	// Mentionned someone?
	if len(m.Mentions) == 1 {

		// Mentionned me?
		if m.Mentions[0].ID == Me.ID {

			// Split
			command := strings.Split(m.Content, " ")

			// Redirect commands
			if len(command) > 1 {
				switch command[1] {
				case "prune":
					commands.Prune(s, guild, channel, m.Message)
					return
				case "get":
					commands.Get(Master, DB, s, guild, channel, m.Message, command)
					return
				case "set":
					commands.Set(DB, s, guild, channel, m.Message, command)
					return
				}
			}
		}
	}

	// Love!
	done = love(s, guild, channel, m.Message)
	if done {
		return
	}
}

func reactHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get the message structure
	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		fmt.Println("Couldn't get the message structure of a MessageReactionAdd!")
		fmt.Println("ChannelID : " + m.ChannelID)
		fmt.Println(err.Error())
		return
	}

	// Get channel structure
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a MessageReactionAdd!")
		fmt.Println("ChannelID : " + m.ChannelID)
		fmt.Println("Author : " + message.Author.Username)
		fmt.Println("Message : " + message.Content)
		fmt.Println(err.Error())
		return
	}

	// Get the guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild structure of a MessageReactionAdd!")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + message.Author.Username)
		fmt.Println("Message : " + message.Content)
		fmt.Println(err.Error())
		return
	}

	// Pin popular message
	pin(s, guild, channel, message)
}

func leaveHandler(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild of " + m.User.Username + "!")
		fmt.Println(err.Error())
		return
	}

	// Invite people who leave
	waitComeBack(s, guild, m.Member)
}

func joinHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild", m.User.Username, "just joined.")
		fmt.Println(err.Error())
		return
	}

	// Ask for guard
	askForGuard(s, guild, m.Member)
}
