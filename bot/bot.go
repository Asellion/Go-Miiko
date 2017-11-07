package bot

import (
	"fmt"
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

	// Ask for guard
	if m.Type == discordgo.MessageTypeGuildMemberJoin {
		askForGuard(s, m)
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
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), "aucun") {
		gardes = append(gardes, "PNJ")
	}

	// Check if there's only one mentionned role
	var garde string
	if len(gardes) == 1 {
		garde = gardes[0]
	}

	// Announce
	if garde == "Étincelante" {
		roleID := getRoleByName(s, channel.GuildID, garde)
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+roleID+">, envoie un message à <@"+guild.OwnerID+"> sur Eldarya pour annoncer ta présence.")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if garde == "Obsidienne" || garde == "Absynthe" || garde == "Ombre" {
		roleID := getRoleByName(s, channel.GuildID, garde)
		s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		_, err := s.ChannelMessageSend(m.ChannelID, "Bienvenue dans la Garde <@&"+roleID+">!")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func askForGuard(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ask newcomer what's their guard
	_, err := s.ChannelMessageSend(m.ChannelID, "Bonjour <@"+m.Author.ID+">! T'es dans quelle garde?")
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
