package bot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
)

// Ask for the guard.
func askForGuard(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	welcomeChannelID := config.GetWelcomeChannelByGuildID(m.GuildID)
	if welcomeChannelID == "" {
		fmt.Println("There are no defined welcome channel for this guild.")
		return
	}

	// Typing!
	err := s.ChannelTyping(welcomeChannelID)
	if err != nil {
		fmt.Println(err.Error())
	}

	if !m.User.Bot {

		// Ask newcomer what's their guard
		_, err = s.ChannelMessageSend(welcomeChannelID, getWelcomeMessage(m.User.ID))
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {

		// Fear the bot!
		_, err = s.ChannelMessageSend(welcomeChannelID, getWelcomeBotMessage(m.User.ID))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getWelcomeMessage(username string) string {

	// Welcome!
	var welcomeList []string
	welcomeList = append(welcomeList, "Bonjour <@"+username+">!")
	welcomeList = append(welcomeList, "Bonjour, <@"+username+">.")
	welcomeList = append(welcomeList, "Bienvenue, <@"+username+">!")
	welcomeList = append(welcomeList, "Bienvenue, <@"+username+">.")
	welcomeList = append(welcomeList, "Bienvenue à <@"+username+">!")
	welcomeList = append(welcomeList, "Bienvenue à toi, <@"+username+">.")
	welcomeList = append(welcomeList, "Bienvenue dans notre serveur, <@"+username+">!")
	welcomeList = append(welcomeList, "Bienvenue dans notre serveur, <@"+username+">.")
	welcomeList = append(welcomeList, "Salutations, <@"+username+">.")
	welcomeList = append(welcomeList, "Ah, <@"+username+">! Nous t'attendions.")
	welcomeList = append(welcomeList, "<@"+username+">, tu es là! Je te souhaite la bienvenue.")
	welcomeList = append(welcomeList, "<@"+username+">, tu es là! Je te souhaite la bienvenue sur notre serveur.")
	welcomeList = append(welcomeList, "<@"+username+">, tu es là! Nous t'attendions.")
	welcomeList = append(welcomeList, "Ah, voilà <@"+username+">. Bienvenue!")
	welcomeList = append(welcomeList, "Ah, voilà <@"+username+">. Je te souhaite la bienvenue!")
	welcomeList = append(welcomeList, "Ah, voilà <@"+username+">. Je te souhaite la bienvenue sur notre serveur.")
	welcomeList = append(welcomeList, "Ah, voilà <@"+username+">. Nous t'attendions.")
	welcomeList = append(welcomeList, "<@"+username+">, je te souhaite la bienvenue.")
	welcomeList = append(welcomeList, "<@"+username+">! Je te souhaite la bienvenue.")
	welcomeList = append(welcomeList, "<@"+username+">, je te souhaite la bienvenue sur notre serveur.")
	welcomeList = append(welcomeList, "<@"+username+">, nous t'attendions.")
	welcomeList = append(welcomeList, "Je te souhaite la bienvenue, <@"+username+">.")
	welcomeList = append(welcomeList, "Je te souhaite la bienvenue, <@"+username+">!")
	welcomeList = append(welcomeList, "Je te souhaite la bienvenue sur notre serveur, <@"+username+">.")
	welcomeList = append(welcomeList, "Nous t'attendions, <@"+username+">.")
	welcomeList = append(welcomeList, "J'ai le plaisir de vous présenter le nouveau membre du serveur, <@"+username+">!")
	welcomeList = append(welcomeList, "J'ai le plaisir de vous présenter le nouveau membre du quartier général, <@"+username+">!")
	welcomeList = append(welcomeList, "Souhaitez tous la bienvenue à <@"+username+">!")
	welcomeList = append(welcomeList, "Une bonne main d'applaudissement pour <@"+username+">!")

	// What's your guard?
	var questionList []string
	questionList = append(questionList, "Dans quelle garde es-tu?")
	questionList = append(questionList, "Quelle est ta garde?")
	questionList = append(questionList, "De quelle garde fais-tu partie?")
	questionList = append(questionList, "Peux-tu me dire tu es dans quelle garde?")
	questionList = append(questionList, "Peux-tu me dire quelle est ta garde?")
	questionList = append(questionList, "Peux-tu me dire de quelle garde tu fais partie?")
	questionList = append(questionList, "Dis-moi, dans quelle garde es-tu?")
	questionList = append(questionList, "Dis-moi, quelle est ta garde?")
	questionList = append(questionList, "Dis-moi, de quelle garde fais-tu partie?")
	questionList = append(questionList, "D'ailleurs, dans quelle garde es-tu?")
	questionList = append(questionList, "D'ailleurs, quelle est ta garde?")
	questionList = append(questionList, "D'ailleurs, de quelle garde fais-tu partie?")
	questionList = append(questionList, "Alors, dans quelle garde es-tu?")
	questionList = append(questionList, "Alors, quelle est ta garde?")
	questionList = append(questionList, "Alors, de quelle garde fais-tu partie?")

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return welcomeList[rand.Intn(len(welcomeList))] + " " + questionList[rand.Intn(len(questionList))]
}

func getWelcomeBotMessage(userID string) string {

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Welcome!
	var welcomeBotList []string

	// Wait, what?
	welcomeBotList = append(welcomeBotList, "Mais... <@"+userID+"> est un bot! Qu'est-ce cette chose fait ici?")
	welcomeBotList = append(welcomeBotList, "Mais quel genre de Faery est <@"+userID+">?")

	// Nope.
	welcomeBotList = append(welcomeBotList, "Non, <@"+userID+">. Je ne veux pas te voir ici.")
	welcomeBotList = append(welcomeBotList, "Hé, <@"+userID+">. On ne veut pas de toi ici.")
	welcomeBotList = append(welcomeBotList, "Arrière, <@"+userID+">!")

	// Botpocalypse
	welcomeBotList = append(welcomeBotList, "T'es venu prendre mon job, <@"+userID+">?")

	// Passive roast
	welcomeBotList = append(welcomeBotList, "Ça pue, ici! Oh, c'est juste <@"+userID+">.")
	welcomeBotList = append(welcomeBotList, "Qui vote pour qu'on kick <@"+userID+">?")
	welcomeBotList = append(welcomeBotList, "On accueille les déchets, maintenant?")
	welcomeBotList = append(welcomeBotList, "Mais quelle abomination!")
	welcomeBotList = append(welcomeBotList, "Beurk.")

	// Notice me senpai!
	welcomeBotList = append(welcomeBotList, "Tiens, un truc moche.")
	welcomeBotList = append(welcomeBotList, "Tiens, un tas de ferraille.")
	welcomeBotList = append(welcomeBotList, "Oh, ça, c'est pas joli.")

	// Community
	welcomeBotList = append(welcomeBotList, "100 PO à celui qui débranche <@"+userID+">!")

	return welcomeBotList[rand.Intn(len(welcomeBotList))]
}

func placeInAGuard(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get a channel structure.")
		fmt.Println("Author : " + m.Author.Username)
		fmt.Println("Message : " + m.Content)
		fmt.Println(err.Error())
	}

	// Get guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get a guild structure.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println(err.Error())
	}

	// Get guild member
	member, err := s.GuildMember(channel.GuildID, m.Author.ID)
	if err != nil {
		fmt.Println("Couldn't get a member structure.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println("Author : " + m.Author.Username)
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
	if strings.Contains(strings.ToLower(m.Content), "eel") || strings.Contains(strings.ToLower(m.Content), "aucun") || strings.Contains(strings.ToLower(m.Content), "ai pas") || strings.Contains(strings.ToLower(m.Content), "pas encore") {
		gardes = append(gardes, "Eel")
	}
	if strings.Contains(strings.ToLower(m.Content), "joue pas") || strings.Contains(strings.ToLower(m.Content), " quoi") {
		gardes = append(gardes, "PNJ")
	}

	// Check if there's only one mentionned role
	var garde string
	if len(gardes) == 1 {
		garde = gardes[0]
	} else {
		return
	}

	// Typing!
	err = s.ChannelTyping(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't tell that I'm typing.")
		fmt.Println("Channel : " + channel.Name)
		fmt.Println(err.Error())
	}

	// Announce
	roleID := getRoleByName(s, channel.GuildID, garde)

	if garde == "Étincelante" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Si tu fais partie de la Garde <@&"+roleID+">, envoie un message à <@"+guild.OwnerID+"> sur Eldarya pour annoncer ta présence. En attendant, dans quelle garde est ton personnage sur Eldarya?")
		if err != nil {
			fmt.Println("Couldn't send message for special role.")
			fmt.Println("Channel : " + channel.Name)
			fmt.Println(err.Error())
		}
		return
	}

	if garde == "Obsidienne" || garde == "Absynthe" || garde == "Ombre" {

		// Add role
		err := s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, getGuardMessage(m.Author.ID, roleID))
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + channel.Name)
			fmt.Println(err.Error())
		}

		return
	}

	if garde == "Eel" {

		// Add role
		err := s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+m.Author.ID+">. Je t'ai donné le rôle <@&"+roleID+"> en attendant que tu rejoignes une garde.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + channel.Name)
			fmt.Println(err.Error())
		}

		return
	}

	if garde == "PNJ" {

		// Add role
		err := s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, roleID)
		if err != nil {
			fmt.Println("Couldn't add a role.")
			fmt.Println("Guild : " + guild.Name)
			fmt.Println("Member : " + m.Author.Username)
			fmt.Println(err.Error())
			return
		}

		// Announce
		_, err = s.ChannelMessageSend(m.ChannelID, "D'accord, <@"+m.Author.ID+">. Je t'ai donné le rôle <@&"+roleID+">, mais saches que ce serveur est dédié à Eldarya.")
		if err != nil {
			fmt.Println("Couldn't announce new role.")
			fmt.Println("Channel : " + channel.Name)
			fmt.Println(err.Error())
		}

		return
	}
}

func getRoleByName(s *discordgo.Session, guildID string, name string) string {

	// Get roles
	guildRoles, err := s.GuildRoles(guildID)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		// Get the first occurence
		for x := 0; x < len(guildRoles); x++ {
			if guildRoles[x].Name == name {
				return guildRoles[x].ID
			}
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
		return ""
	}

	// Create the missing role
	role, err := s.GuildRoleCreate(guildID)
	if err != nil {
		fmt.Println("Couldn't create a role.")
		fmt.Println(err.Error())
		return ""
	}

	// Edit the missing role
	_, err = s.GuildRoleEdit(guildID, role.ID, name, color, false, role.Permissions, false)
	if err != nil {
		fmt.Println("Couldn't edit the permissions of the newly created role " + name + ".")
		fmt.Println(err.Error())
		return ""
	}

	return role.ID
}

func getGuardMessage(userID string, roleID string) string {

	// Messages
	var messageList []string
	messageList = append(messageList, "Bienvenue à <@"+userID+"> dans la garde <@&"+roleID+">!")
	messageList = append(messageList, "Bienvenue dans la garde <@&"+roleID+">, <@"+userID+">.")
	messageList = append(messageList, "Bienvenue dans la garde <@&"+roleID+">, <@"+userID+">. J'espère que tu ne t'attends pas à un matelas.")
	messageList = append(messageList, "Bienvenue au sein la garde <@&"+roleID+">, <@"+userID+">.")
	messageList = append(messageList, "Bienvenue parmis les <@&"+roleID+">, <@"+userID+">.")
	messageList = append(messageList, "<@"+userID+"> est maintenant un membre de la garde <@&"+roleID+">!")
	messageList = append(messageList, "<@&"+roleID+"> a l'honneur d'accueillir <@"+userID+">!")
	messageList = append(messageList, "<@&"+roleID+">! Faites de la place pour <@"+userID+">!")
	messageList = append(messageList, "<@"+userID+"> fait maintenant partie de la garde <@&"+roleID+">.")
	messageList = append(messageList, "<@"+userID+"> est maintenant une <@&"+roleID+">!")
	messageList = append(messageList, "Souhaitez la bienvenue à notre nouvelle <@&"+roleID+">, <@"+userID+">!")
	messageList = append(messageList, "Bien! <@"+userID+"> a maintenant une place dans les cachots de la garde <@&"+roleID+">.")
	messageList = append(messageList, "<@"+userID+"> a rejoint la garde <@&"+roleID+">.")
	messageList = append(messageList, "Je savais que <@"+userID+"> était une <@&"+roleID+">!")
	messageList = append(messageList, "Ah, je savais que <@"+userID+"> était une <@&"+roleID+">.")
	messageList = append(messageList, "Je savais bien que <@"+userID+"> était une <@&"+roleID+">!")
	messageList = append(messageList, "Ah, je le savais! <@"+userID+"> est une <@&"+roleID+">!")
	messageList = append(messageList, "J'en étais sûre! <@"+userID+"> est une <@&"+roleID+">!")
	messageList = append(messageList, "<@"+userID+"> est dorénavant une <@&"+roleID+">.")
	messageList = append(messageList, "Accueillez notre nouvelle <@&"+roleID+">, <@"+userID+">!")
	messageList = append(messageList, "Je te souhaite un bon séjour parmis les <@&"+roleID+">, <@"+userID+">.")
	messageList = append(messageList, "<@"+userID+"> peut maintenant rejoindre les <@&"+roleID+">.")
	messageList = append(messageList, "Tu peux rejoindre les <@&"+roleID+">, <@"+userID+">.")
	messageList = append(messageList, "Que les <@&"+roleID+"> soient avec <@"+userID+">.")

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return messageList[seed.Intn(len(messageList))]
}
