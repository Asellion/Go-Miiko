package bot

import (
	"fmt"
	"math"
	"strings"

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
			fmt.Println(m.Author.Username + " : " + m.Content)
		} else if m.ChannelID == config.BotMasterChannelID {
		} else {
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

func placeInAGuard(s *discordgo.Session, m *discordgo.MessageCreate) {

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

	// Get guild member
	member, err := s.GuildMember(channel.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	// If Author has no role
	if len(member.Roles) != 0 {
		return
	}

	// Get mentionned roles
	var gardes []string
	if strings.Contains(strings.ToLower(m.Content), "tincelant") {
		gardes = append(gardes, "Étincelante")
	}
	if strings.Contains(strings.ToLower(m.Content), "obsi") {
		gardes = append(gardes, "Obsidienne")
	}
	if strings.Contains(strings.ToLower(m.Content), "absy") {
		gardes = append(gardes, "Absynthe")
	}
	if strings.Contains(strings.ToLower(m.Content), "ombr") {
		gardes = append(gardes, "Ombre")
	}
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), "aucun") || strings.Contains(strings.ToLower(m.Content), "ai pas") || strings.Contains(strings.ToLower(m.Content), " quoi") {
		gardes = append(gardes, "PNJ")
	}

	// Check if there's only one mentionned role
	var garde string
	if len(gardes) == 1 {
		garde = gardes[0]
	}

	// Announce
	roleID := getRoleByName(s, channel.GuildID, garde)
	if garde == "Étincelante" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+roleID+">, envoie un message à <@"+guild.OwnerID+"> sur Eldarya pour annoncer ta présence. En attendant, dans quelle garde est ton personnage sur Eldarya?")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if garde == "Obsidienne" || garde == "Absynthe" || garde == "Ombre" {
		s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		_, err := s.ChannelMessageSend(m.ChannelID, "Bienvenue à <@"+m.Author.ID+"> dans la Garde <@&"+roleID+">!")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if garde == "PNJ" {
		s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		_, err := s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+m.Author.ID+">. Je t'ai donné le rôle <@&"+roleID+"> en attendant que tu rejoignes une garde.")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func askForGuard(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ask newcomer what's their guard
	_, err := s.ChannelMessageSend(m.ChannelID, "Bonjour <@"+m.Author.ID+">! De quelle garde fais-tu partie?")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getRoleByName(s *discordgo.Session, guildID string, name string) string {

	// Get roles
	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get the first occurence
	for x := 0; x < len(guildRoles); x++ {
		if guildRoles[x].Name == name {
			return guildRoles[x].ID
		}
	}

	return ""
}

func reactHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get channel structure
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a MessageReactionAdd!")
		fmt.Println("m.ChannelID : " + m.ChannelID)
		fmt.Println(err.Error())
	}

	// DM?
	if channel.Type == discordgo.ChannelTypeDM {
		return
	}

	// Get the message structure
	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		fmt.Println("Couldn't get the message structure of a MessageReactionAdd!")
		fmt.Println("m.ChannelID : " + m.ChannelID)
		fmt.Println("m.MessageID : " + m.MessageID)
		fmt.Println(err.Error())
	}

	// Get the guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild structure of a MessageReactionAdd!")
		fmt.Println(err.Error())
	}

	// Get phi
	phi := (1 + math.Sqrt(5)) / 2

	// Get people online
	var onlineCount int
	for x := 0; x < len(guild.Presences); x++ {
		if guild.Presences[x].Status == discordgo.StatusOnline {
			onlineCount++
		}
	}

	// Get online / phi
	min := int(float64(len(guild.Presences)) / phi)

	// Count the reactions
	if len(message.Reactions) > min {

		// Pin it!
		err := s.ChannelMessagePin(m.ChannelID, m.MessageID)
		if err != nil {
			fmt.Println("Couldn't pin a popular message!")
			fmt.Println(err.Error())
		}
	}
}
