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
	err := DB.QueryRow("select `channel` from `welcome` where `server` = ?", g.ID).Scan(&cid)
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
		_, err = s.ChannelMessageSend(channel.ID, getByeBotMessage(m.User))
		if err != nil {
			fmt.Println("Couldn't say bye to " + m.User.Username + "!")
			fmt.Println(err.Error())
		}

	} else {

		// Announce departure
		err = s.ChannelTyping(channel.ID)
		_, err = s.ChannelMessageSend(channel.ID, getPublicByeMessage(m.User))
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

func getPublicByeMessage(user *discordgo.User) string {

	// Bye Messages
	byeList := [...]string{

		// Messages
		"J'ai le regret d'annoncer le départ de **" + user.Username + "**.",
		"C'est avec émotion que j'annonce le départ de **" + user.Username + "**.",
		"L'Oracle a emporté **" + user.Username + "** avec elle.",
		"**" + user.Username + "** a quitté la garde.",
		"Attends, **" + user.Username + "**, reviens!",
		"**" + user.Username + "** a pris son envol!",
		"**" + user.Username + "** vole de ses propres ailes.",
		"**" + user.Username + "** part à l'aventure!",
		"L'aventure de **" + user.Username + "** se termine ici.",
		"La garde se souviendra de **" + user.Username + "**.",
		"Il pleut lorsque je regarde vers **" + user.Username + "**.",
		"Mon coeur se serre à l'annonce du départ de **" + user.Username + "**.",
		"**" + user.Username + "** a donné sa démission.",
		"Que la force soit avec **" + user.Username + "**.",

		// Death
		"Repose en paix, **" + user.Username + "**.",
		"Pourquoi, **" + user.Username + "**, pourquoi?",
		"**" + user.Username + "** s'est fait dévorer par un Hydracarys.",
		"**" + user.Username + "** a marché dans une toile de Chead.",
		"Un Black Gallytrot a démembré **" + user.Username + "**.",
		"La foudre a frappé **" + user.Username + "**.",
		"Je suis attristée d'apprendre la perte soudaine de **" + user.Username + "**.",
		"Mon coeur est avec **" + user.Username + "** en ce temps de tristesse.",
		"Mes sincères condoléances pour la perte de **" + user.Username + "**.",
		"Les mots ne peuvent exprimer combien je suis attristée d'apprendre la perte de **" + user.Username + "**.",
		"Mes pensées et mes prières sont avec **" + user.Username + "** pendant cette période tragique.",
		"Mes plus sincères condoléances pour la perte de **" + user.Username + "**.",
		"Que Dieu bénisse **" + user.Username + "** en ce moment de tristesse.",
		"Je suis vraiment attristée d'apprendre la mort de **" + user.Username + "**.",
		"Puisse **" + user.Username + "** reposer en paix.",
		"C'est avec une grande tristesse que j'ai appris le décès de **" + user.Username + "**.",

		// Community
		"Aurevoir, **" + user.Username + "**. Reviens-nous vite!",
		"**" + user.Username + "** nous a quitté. Souhaiton-lui le meilleur!",
		"**" + user.Username + "** nous a quitté. Elle va nous manquer.",
		"Adieu, **" + user.Username + "**! Vole vers d'autres cieux!",
		"**" + user.Username + "** a été transféré vers un autre QG.",
		"Nous n'oublierons pas le sacrifice de **" + user.Username + "**!",
		"Nous avons perdu **" + user.Username + "**, mais nous restons forts.",

		// Legendary
		"C'est en ce jour funeste que **" + user.Username + "** nous a quitté. Puisse son âme rejoindre le cristal et son héritage mon porte-maanas.",
	}

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeList[rand.Intn(len(byeList))]
}

func getByeBotMessage(user *discordgo.User) string {

	// Bye Messages
	byeBotList := [...]string{

		// Messages
		"Bon débarras, **" + user.Username + "**.",
		"Bien! Personne ne va s'ennuyer de **" + user.Username + "**.",
		"De toute façon, **" + user.Username + "** n'avait aucun lien avec Eldarya.",
		"**" + user.Username + "** ne nous manquera pas.",
		"Ha! **" + user.Username + "** est parti. Ça fait plus de popcorn pour moi!",

		// Community
		"Nous sommes enfin débarrassés de **" + user.Username + "**!",
		"Oh, **" + user.Username + "** est mort. Mais quel dommage.",
		"Super! **" + user.Username + "** a fiché le camp!",
		"Ah? **" + user.Username + "** était là?",
	}

	// Random
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return byeBotList[rand.Intn(len(byeBotList))]
}
