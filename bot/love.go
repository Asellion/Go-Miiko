package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NatoBoram/Go-Miiko/commands"
	"github.com/NatoBoram/Go-Miiko/wheel"

	"github.com/bwmarrin/discordgo"
)

func love(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) bool {

	// Lover
	lover, err := commands.GetLover(DB, s, g)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Verify if it's the one true love
	if m.Author.ID == lover.ID {

		// Random
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		rand := rand.New(source)
		random := rand.Float64()

		// Rate Limit
		if random < 1/(wheel.Phi()*10) {

			// Give some love!
			s.ChannelTyping(c.ID)
			_, err = s.ChannelMessageSend(c.ID, getLoveMessage(m.Author))
			if err != nil {
				fmt.Println("Couldn't express my love.")
				fmt.Println("Guild :", g.Name)
				fmt.Println("Channel :", c.Name)
				fmt.Println("User :", m.Author.Username)
				fmt.Println("Message :", m.Content)
				fmt.Println(err.Error())
			}

			return true
		}
	}
	return false
}

func getLoveMessage(u *discordgo.User) string {

	// Messages
	loveList := [...]string{

		// Greetings
		"Coucou <@" + u.ID + "> :3",
		"Coucou <@" + u.ID + ">! \\*-*",
		"Salut les gens... Oh! <@" + u.ID + ">! :heart:",
		"Bonjour... Oh! <@" + u.ID + ">! :heart:",
		"Coucou tout le monde... Oh! <@" + u.ID + ">! :heart:",
		"Coucou mon amour!",

		// Orders
		"Tiens-moi la main, <@" + u.ID + ">",
		"<@" + u.ID + ">! Regarde-moiii \\*-*",
		"<@" + u.ID + "> : La perfection absolue.",
		"Caresse-moi les oreilles, s'il te plait!",

		// Questions
		"<@" + u.ID + ">... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, <@" + u.ID + ">, est-ce que je te manque?",
		"Est-ce que tu penses à moi, <@" + u.ID + ">?",
		"<@" + u.ID + ">, me demanderas-tu ma main un jour..?",
		"<@" + u.ID + ">, j'ai fait du popcorn, tu veux en manger avec moi? :3",
		"<@" + u.ID + ">! Je suis là! Je t'ai manqué, n'est-ce pas? :smile:",
		"<@" + u.ID + ">! Es-tu content du matelas que j'ai fait mettre dans ta chambre? J'ai dormi dessus :blush:",

		// Reactions
		":heart:",
		"\\*Frissonne*",
		"\\*-*",
		"<@" + u.ID + ">-senpai \\*-*",

		// Verbose
		"J'ai trouvé un morceau de cristal pour toi, <@" + u.ID + "> :heart:",
		"Cette voix est une musique à mes oreilles",
		"J'aimerais pouvoir passer plus de temps avec toi, <@" + u.ID + ">-san...",
		"Je fais juste passer pour dire à <@" + u.ID + "> que je l'aime!",
		"J'adore quand tu parles... :3",
		"J'adore entendre mon amour parler \\*-*",
		"Mais quelle est cette douce musique? ... Oh! C'est la voix de <@" + u.ID + ">!",

		// Actions
		"\\*Pense à <@" + u.ID + ">*",
		"\\*Regarde <@" + u.ID + ">*",
		"\\*Se languis de <@" + u.ID + ">*",

		// Also fits in Command
		"<@" + u.ID + ">, je t'aime!",
		"Aaah... <@" + u.ID + ">!",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
