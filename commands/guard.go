package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// PlaceInAGuard gives members a role.
func PlaceInAGuard(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.Member, m *discordgo.Message) {

	// If Author has no role
	if len(u.Roles) != 0 {
		return
	}

	// Get mentionned roles
	gardes := getMentionnedGuard(m)

	// Check if there's only one mentionned role
	var garde string
	if len(gardes) == 1 {
		garde = gardes[0]
	} else if len(gardes) != 0 {
		// Sorry, I didn't understand.
		return
	} else {
		// Why are you ignoring me?
		return
	}

	// Typing!
	err := s.ChannelTyping(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't tell that I'm typing.")
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
	}

	// Announce
	role := getRoleByName(s, g, garde)

	if garde == "Étincelante" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+role.ID+">, envoie un message à <@"+g.OwnerID+"> sur Eldarya pour annoncer ta présence. En attendant, dans quelle garde est ton personnage sur Eldarya?")
		if err != nil {
			fmt.Println("Couldn't send message for special role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}
		return
	}

	if garde == "Obsidienne" || garde == "Absynthe" || garde == "Ombre" {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, role.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Role : " + role.ID)
			fmt.Println("Member : " + u.User.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, getGuardMessage(u.User, role))
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}

		return
	}

	if garde == "Eel" {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, role.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+role.ID+"> en attendant que tu rejoignes une garde.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}

		return
	}

	if garde == "PNJ" {

		// Add role
		err := s.GuildMemberRoleAdd(g.ID, u.User.ID, role.ID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+u.User.ID+">. Je t'ai donné le rôle <@&"+role.ID+">, mais saches que ce serveur est dédié à Eldarya.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + c.Name)
			fmt.Println(err.Error())
		}

		return
	}
}

func getMentionnedGuard(m *discordgo.Message) []string {
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
	if strings.Contains(strings.ToLower(m.Content), "eel") || strings.Contains(strings.ToLower(m.Content), "aucun") || strings.Contains(strings.ToLower(m.Content), "ai pas") || strings.Contains(strings.ToLower(m.Content), "pas encore") {
		gardes = append(gardes, "Eel")
	}
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), " quoi") {
		gardes = append(gardes, "PNJ")
	}
	return gardes
}

func getRoleByName(s *discordgo.Session, g *discordgo.Guild, name string) *discordgo.Role {

	// Get roles
	guildRoles, err := s.GuildRoles(g.ID)
	if err != nil {
		fmt.Println("Couldn't get the guild's roles.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Role :", name)
		fmt.Println(err.Error())
	}

	// Get the first occurrence
	for x := 0; x < len(guildRoles); x++ {
		if guildRoles[x].Name == name {
			return guildRoles[x]
		}
	}

	// Get color
	var color int
	if name == "Étincelante" {
		color = 16705182
	} else if name == "Obsidienne" {
		color = 16496296
	} else if name == "Absynthe" {
		color = 8321915
	} else if name == "Ombre" {
		color = 16364540
	} else if name == "Eel" {
		color = 12503544
	} else if name == "PNJ" {
		color = 10263708
	} else {
		return nil
	}

	// Create the missing role
	role, err := s.GuildRoleCreate(g.ID)
	if err != nil {
		fmt.Println("Couldn't create a role.")
		fmt.Println(err.Error())
		return nil
	}

	// Edit the missing role
	_, err = s.GuildRoleEdit(g.ID, role.ID, name, color, false, role.Permissions, false)
	if err != nil {
		fmt.Println("Couldn't edit the newly created role,", name+".")
		fmt.Println(err.Error())
		return nil
	}

	return role
}

func getGuardMessage(user *discordgo.User, role *discordgo.Role) string {

	// Messages
	messageList := [...]string{
		"Bienvenue à <@" + user.ID + "> dans la garde <@&" + role.ID + ">!",
		"Bienvenue dans la garde <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Bienvenue dans la garde <@&" + role.ID + ">, <@" + user.ID + ">. J'espère que tu ne t'attends pas à un matelas.",
		"Bienvenue au sein la garde <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Bienvenue parmis les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"<@" + user.ID + "> est maintenant un membre de la garde <@&" + role.ID + ">!",
		"<@&" + role.ID + "> a l'honneur d'accueillir <@" + user.ID + ">!",
		"<@&" + role.ID + ">! Faites de la place pour <@" + user.ID + ">!",
		"<@" + user.ID + "> fait maintenant partie de la garde <@&" + role.ID + ">.",
		"<@" + user.ID + "> est maintenant une <@&" + role.ID + ">!",
		"Souhaitez la bienvenue à notre nouvelle <@&" + role.ID + ">, <@" + user.ID + ">!",
		"Bien! <@" + user.ID + "> a maintenant une place dans les cachots de la garde <@&" + role.ID + ">.",
		"<@" + user.ID + "> a rejoint la garde <@&" + role.ID + ">.",
		"Je savais que <@" + user.ID + "> était une <@&" + role.ID + ">!",
		"Ah, je savais que <@" + user.ID + "> était une <@&" + role.ID + ">.",
		"Je savais bien que <@" + user.ID + "> était une <@&" + role.ID + ">!",
		"Ah, je le savais! <@" + user.ID + "> est une <@&" + role.ID + ">!",
		"J'en étais sûre! <@" + user.ID + "> est une <@&" + role.ID + ">!",
		"<@" + user.ID + "> est dorénavant une <@&" + role.ID + ">.",
		"Accueillez notre nouvelle <@&" + role.ID + ">, <@" + user.ID + ">!",
		"Je te souhaite un bon séjour parmis les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"<@" + user.ID + "> peut maintenant rejoindre les <@&" + role.ID + ">.",
		"Tu peux rejoindre les <@&" + role.ID + ">, <@" + user.ID + ">.",
		"Que les <@&" + role.ID + "> soient avec <@" + user.ID + ">.",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return messageList[seed.Intn(len(messageList))]
}
