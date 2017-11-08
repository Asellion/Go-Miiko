package bot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Ask for the guard.
func askForGuard(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ask newcomer what's their guard
	_, err := s.ChannelMessageSend(m.ChannelID, getWelcomeMessage(m.Author.ID))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getWelcomeMessage(username string) string {

	// Welcome!
	var welcomeList []string
	welcomeList = append(welcomeList, "Bonjour <@"+username+">!")
	welcomeList = append(welcomeList, "Bonjour, <@"+username+">.")
	welcomeList = append(welcomeList, "Salut <@"+username+">!")
	welcomeList = append(welcomeList, "Bienvenue dans notre serveur, <@"+username+">.")

	// What's your guard?
	var questionList []string
	questionList = append(questionList, "Tu es dans quelle garde?")
	questionList = append(questionList, "Dans quelle garde es-tu?")
	questionList = append(questionList, "Tu fais partie de quelle garde?")

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return welcomeList[seed.Intn(len(welcomeList))] + " " + questionList[seed.Intn(len(questionList))]
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
