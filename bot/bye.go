package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func waitComeBack(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Member) {

	// Get welcome channel
	var cid string
	err := DB.QueryRow("select `welcome` from `servers` where `server` = ?", g.ID).Scan(&cid)
	if err != nil {
		fmt.Println("Couldn't select the welcome channel of", g.Name, ".")
		return
	}

	// Make sure the channel exists
	channel, err := s.Channel(cid)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a welcome channel.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("ChannelID : " + cid)
		fmt.Println(err.Error())
		return
	}

	// Bot?
	if m.User.ID == Me.ID {
		fmt.Println("Looks like I just left", g.Name+".")
	} else if m.User.Bot {

		// Bye bot!
		s.ChannelTyping(channel.ID)
		_, err = s.ChannelMessageSend(channel.ID, getByeBotMessage(m.User.ID))
		if err != nil {
			fmt.Println("Couldn't say bye to " + m.User.Username + "!")
			fmt.Println(err.Error())
		}

	} else {

		// Announce departure
		err = s.ChannelTyping(channel.ID)
		_, err = s.ChannelMessageSend(channel.ID, getPublicByeMessage(m.User.ID))
		if err != nil {
			fmt.Println("Couldn't announce the departure of " + m.User.Username + ".")
			fmt.Println(err.Error())
		}

		// Create an invite structure
		var invStruct discordgo.Invite
		invStruct.Temporary = true

		// Create an invite to WelcomeChannel
		var invite *discordgo.Invite
		invite, err = s.ChannelInviteCreate(channel.ID, invStruct)
		if err != nil {
			fmt.Println("Couldn't create an invite in " + g.Name + ".")
			fmt.Println(err.Error())
			return
		}

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
			fmt.Println("Couldn't send a message to " + m.User.Username + "!")
			fmt.Println(err.Error())
		}
	}
}

func getPrivateByeMessage(inviteCode string) string {

	// Bye Messages
	byeList := [...]string{

		// Messages
		"Oh, je suis triste de te voir partir! Si tu veux nous rejoindre à nouveau, j'ai créé une invitation pour toi : https://discord.gg/" + inviteCode,
		"Au revoir! Voici une invitation au cas où tu changes d'idée : https://discord.gg/" + inviteCode,
		"Tu vas me manquer. Si tu veux me revoir, j'ai créé une invitation pour toi : https://discord.gg/" + inviteCode,
	}

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getPublicByeMessage(userID string) string {

	// Bye Messages
	byeList := [...]string{

		// Messages
		"J'ai le regret d'annoncer le départ de <@" + userID + ">.",
		"C'est avec émotion que j'annonce le départ de <@" + userID + ">.",
		"L'Oracle a emporté <@" + userID + "> avec elle.",
		"<@" + userID + "> a quitté la garde.",
		"Attends, <@" + userID + ">, reviens!",
		"<@" + userID + "> a pris son envol!",
		"<@" + userID + "> vole de ses propres ailes.",
		"<@" + userID + "> part à l'aventure!",
		"L'aventure de <@" + userID + "> se termine ici.",
		"La garde se souviendra de <@" + userID + ">.",
		"Il pleut lorsque je regarde vers <@" + userID + ">.",
		"Mon coeur se serre à l'annonce du départ de <@" + userID + ">.",
		"<@" + userID + "> a donné sa démission.",
		"Que la force soit avec <@" + userID + ">.",

		// Death
		"Repose en paix, <@" + userID + ">.",
		"Pourquoi, <@" + userID + ">, pourquoi?",
		"<@" + userID + "> s'est fait dévorer par un Hydracarys.",
		"<@" + userID + "> a marché dans une toile de Chead.",
		"Un Black Gallytrot a démembré <@" + userID + ">.",
		"La foudre a frappé <@" + userID + ">.",
		"Je suis attristée d'apprendre la perte soudaine de <@" + userID + ">.",
		"Mon coeur est avec <@" + userID + "> en ce temps de tristesse.",
		"Mes sincères condoléances pour la perte de <@" + userID + ">.",
		"Les mots ne peuvent exprimer combien je suis attristée d'apprendre la perte de <@" + userID + ">.",
		"Mes pensées et mes prières sont avec <@" + userID + "> pendant cette période tragique.",
		"Mes plus sincères condoléances pour la perte de <@" + userID + ">.",
		"Que Dieu bénisse <@" + userID + "> en ce moment de tristesse.",
		"Je suis vraiment attristée d'apprendre la mort de <@" + userID + ">.",
		"Puisse <@" + userID + "> reposer en paix.",
		"C'est avec une grande tristesse que j'ai appris le décès de <@" + userID + ">.",

		// Community
		"Aurevoir, <@" + userID + ">. Reviens-nous vite!",
		"<@" + userID + "> nous a quitté. Souhaiton-lui le meilleur!",
		"<@" + userID + "> nous a quitté. Elle va nous manquer.",
		"Adieu, <@" + userID + ">! Vole vers d'autres cieux!",
		"<@" + userID + "> a été transféré vers un autre QG.",
		"Nous n'oublierons pas le sacrifice de <@" + userID + ">!",
		"Nous avons perdu <@" + userID + ">, mais nous restons forts.",

		// Legendary
		"C'est en ce jour funeste que <@" + userID + "> nous a quitté. Puisse son âme rejoindre le cristal et son héritage mon porte-maanas.",
	}

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getByeBotMessage(userID string) string {

	// Bye Messages
	byeBotList := [...]string{

		// Messages
		"Bon débarras, <@" + userID + ">.",
		"Bien! Personne ne va s'ennuyer de <@" + userID + ">.",
		"De toute façon, <@" + userID + "> n'avait aucun lien avec Eldarya.",
		"<@" + userID + "> ne nous manquera pas.",
		"Ha! <@" + userID + "> est parti. Ça fait plus de popcorn pour moi!",

		// Community
		"Nous sommes enfin débarrassés de <@" + userID + ">!",
		"Oh, <@" + userID + "> est mort. Mais quel dommage.",
		"Super! <@" + userID + "> a fiché le camp!",
		"Ah? <@" + userID + "> était là?",
	}

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeBotList[rand.Intn(len(byeBotList))]
}
