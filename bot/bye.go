package bot

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func waitComeBack(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	/*
		// Get guild
		guild, err := s.State.Guild(m.GuildID)
		if err != nil {
			fmt.Println("Couldn't get " + m.User.Username + "'s guild ID.")
			fmt.Println(err.Error())
			return
		}

		// Get channel
		channel, exists := config.Database.WelcomeChannels[guild.ID]
		if !exists {
			return
		}

		// Create an invite structure
		var invStruct discordgo.Invite
		invStruct.Temporary = true

		// Create an invite to WelcomeChannel
		var invite *discordgo.Invite
		invite, err = s.ChannelInviteCreate(channel, invStruct)
		if err != nil {
			fmt.Println("Couldn't create an invite in " + guild.Name + ".")
			fmt.Println(err.Error())
			return
		}

		// Bot?
		if m.User.Bot {

			// Bye bot!
			s.ChannelTyping(channel)
			_, err = s.ChannelMessageSend(channel, getByeBotMessage(m.User.ID))
			if err != nil {
				fmt.Println("Couldn't say bye to " + m.User.Username + "!")
				fmt.Println(err.Error())
			}

		} else {

			// Open channel
			privateChannel, err := s.UserChannelCreate(m.User.ID)
			if err != nil {
				fmt.Println("Couldn't create a private channel with " + m.User.Username + ".")
				fmt.Println(err.Error())
				return
			}

			// Send message
			s.ChannelTyping(privateChannel.ID)
			_, err = s.ChannelMessageSend(privateChannel.ID, getPrivateByeMessage(invite.Code))
			if err != nil {
				fmt.Println("Couldn't say bye to " + m.User.Username + "!")
				fmt.Println(err.Error())
			}

			// Announce departure
			err = s.ChannelTyping(channel)
			_, err = s.ChannelMessageSend(channel, getPublicByeMessage(m.User.ID))
			if err != nil {
				fmt.Println("Couldn't announce the departure of " + m.User.Username + ".")
				fmt.Println(err.Error())
			}
		}
	*/
}

func getPrivateByeMessage(inviteCode string) string {

	// Bye Messages
	var byeList []string

	// Messages
	byeList = append(byeList, "Oh, je suis triste de te voir partir! Si tu veux nous rejoindre à nouveau, j'ai créé une invitation pour toi : https://discord.gg/"+inviteCode)
	byeList = append(byeList, "Au revoir! Voici une invitation si tu changes d'idée : https://discord.gg/"+inviteCode)
	byeList = append(byeList, "Tu vas me manquer. Si tu veux me revoir, j'ai créé une invitation pour toi : https://discord.gg/"+inviteCode)

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getPublicByeMessage(userID string) string {

	// Bye Messages
	var byeList []string

	// Messages
	byeList = append(byeList, "J'ai le regret d'annoncer le départ de <@"+userID+">.")
	byeList = append(byeList, "C'est avec émotion que j'annonce le départ de <@"+userID+">.")
	byeList = append(byeList, "L'Oracle a emporté <@"+userID+"> avec elle.")
	byeList = append(byeList, "<@"+userID+"> a quitté la garde.")
	byeList = append(byeList, "Attends, <@"+userID+">, reviens!")
	byeList = append(byeList, "<@"+userID+"> a pris son envol!")
	byeList = append(byeList, "<@"+userID+"> vole de ses propres ailes.")
	byeList = append(byeList, "<@"+userID+"> part à l'aventure!")
	byeList = append(byeList, "L'aventure de <@"+userID+"> se termine ici.")
	byeList = append(byeList, "La garde se souviendra de <@"+userID+">.")
	byeList = append(byeList, "Il pleut lorsque je regarde vers <@"+userID+">.")
	byeList = append(byeList, "Mon coeur se serre à l'annonce du départ de <@"+userID+">.")
	byeList = append(byeList, "<@"+userID+"> a donné sa démission.")
	byeList = append(byeList, "Que la force soit avec <@"+userID+">.")

	// Death
	byeList = append(byeList, "Repose en paix, <@"+userID+">.")
	byeList = append(byeList, "Pourquoi, <@"+userID+">, pourquoi?")
	byeList = append(byeList, "<@"+userID+"> s'est fait dévorer par un Hydracarys.")
	byeList = append(byeList, "<@"+userID+"> a marché dans une toile de Chead.")
	byeList = append(byeList, "Un Black Gallytrot a démembré <@"+userID+">.")
	byeList = append(byeList, "La foudre a frappé <@"+userID+">.")
	byeList = append(byeList, "Je suis attristée d'apprendre la perte soudaine de <@"+userID+">.")
	byeList = append(byeList, "Mon coeur est avec <@"+userID+"> en ce temps de tristesse.")
	byeList = append(byeList, "Mes sincères condoléances pour la perte de <@"+userID+">.")
	byeList = append(byeList, "Les mots ne peuvent exprimer combien je suis attristée d'apprendre la perte de <@"+userID+">.")
	byeList = append(byeList, "Mes pensées et mes prières sont avec <@"+userID+"> pendant cette période tragique.")
	byeList = append(byeList, "Mes plus sincères condoléances pour la perte de <@"+userID+">.")
	byeList = append(byeList, "Que Dieu bénisse <@"+userID+"> en ce moment de tristesse.")
	byeList = append(byeList, "Je suis vraiment attristée d'apprendre la mort de <@"+userID+">.")
	byeList = append(byeList, "Puisse <@"+userID+"> reposer en paix.")
	byeList = append(byeList, "C'est avec une grande tristesse que j'ai appris le décès de <@"+userID+">.")

	// Community
	byeList = append(byeList, "Aurevoir, <@"+userID+">. Reviens-nous vite!")
	byeList = append(byeList, "<@"+userID+"> nous a quitté. Souhaiton-lui le meilleur!")
	byeList = append(byeList, "<@"+userID+"> nous a quitté. Elle va nous manquer.")
	byeList = append(byeList, "Adieu, <@"+userID+">! Vole vers d'autres cieux!")
	byeList = append(byeList, "<@"+userID+"> a été transféré vers un autre QG.")
	byeList = append(byeList, "Nous n'oublierons pas le sacrifice de <@"+userID+">!")
	byeList = append(byeList, "Nous avons perdu <@"+userID+">, mais nous restons forts.")

	// Legendary
	byeList = append(byeList, "C'est en ce jour funeste que <@"+userID+"> nous a quitté. Puisse son âme rejoindre le cristal et son héritage mon porte-maanas.")

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getByeBotMessage(userID string) string {

	// Bye Messages
	var byeBotList []string

	// Messages
	byeBotList = append(byeBotList, "Bon débarras, <@"+userID+">.")
	byeBotList = append(byeBotList, "Bien! Personne ne va s'ennuyer de <@"+userID+">.")
	byeBotList = append(byeBotList, "De toute façon, <@"+userID+"> n'avait aucun lien avec Eldarya.")
	byeBotList = append(byeBotList, "<@"+userID+"> ne nous manquera pas.")
	byeBotList = append(byeBotList, "Ha! <@"+userID+"> est parti. Ça fait plus de popcorn pour moi!")

	// Community
	byeBotList = append(byeBotList, "Nous sommes enfin débarrassés de <@"+userID+">!")
	byeBotList = append(byeBotList, "Oh, <@"+userID+"> est mort. Mais quel dommage.")
	byeBotList = append(byeBotList, "Super! <@"+userID+"> a fiché le camp!")
	byeBotList = append(byeBotList, "Ah? <@"+userID+"> était là?")

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeBotList[rand.Intn(len(byeBotList))]
}
