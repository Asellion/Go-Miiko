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

			// Member
			member, err := s.GuildMember(g.ID, lover.ID)
			if err != nil {
				fmt.Println("Couldn't get the member I love.")
				fmt.Println("Guild :", g.Name)
				fmt.Println("Channel :", c.Name)
				fmt.Println("User :", m.Author.Username)
				fmt.Println("Message :", m.Content)
				fmt.Println(err.Error())
				return false
			}

			mention := "**" + member.Nick + "**"

			// Give some love!
			s.ChannelTyping(c.ID)
			_, err = s.ChannelMessageSend(c.ID, getLoveMessage(mention))
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

func getLoveMessage(name string) string {

	// Messages
	loveList := [...]string{

		// Greetings
		"Coucou " + name + " :3",
		"Coucou " + name + "! \\*-*",
		"Salut les gens... Oh! " + name + "! :heart:",
		"Bonjour... Oh! " + name + "! :heart:",
		"Coucou tout le monde... Oh! " + name + "! :heart:",
		"Coucou mon amour!",

		// Orders
		"Tiens-moi la main, " + name + "",
		"" + name + "! Regarde-moiii \\*-*",
		"" + name + " : La perfection absolue.",
		"Caresse-moi les oreilles, s'il te plait!",

		// Questions
		"" + name + "... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, " + name + ", est-ce que je te manque?",
		"Est-ce que tu penses à moi, " + name + "?",
		"" + name + ", me demanderas-tu ma main un jour..?",
		"" + name + ", j'ai fait du popcorn, tu veux en manger avec moi? :3",
		"" + name + "! Je suis là! Je t'ai manqué, n'est-ce pas? :smile:",
		"" + name + "! Es-tu content du matelas que j'ai fait mettre dans ta chambre? J'ai dormi dessus :blush:",

		// Reactions
		":heart:",
		"\\*Frissonne*",
		"\\*-*",
		"" + name + "-senpai \\*-*",

		// Verbose
		"J'ai trouvé un morceau de cristal pour toi, " + name + " :heart:",
		"Cette voix est une musique à mes oreilles",
		"J'aimerais pouvoir passer plus de temps avec toi, " + name + "...",
		"Je fais juste passer pour dire à " + name + " que je l'aime!",
		"J'adore quand tu parles... :3",
		"J'adore entendre mon amour parler \\*-*",
		"Mais quelle est cette douce musique? ... Oh! C'est la voix de " + name + "!",

		// Actions
		"\\*Pense à " + name + "*",
		"\\*Regarde " + name + "*",
		"\\*Se languis de " + name + "*",

		// Also fits in Command
		"" + name + ", je t'aime!",
		"Aaah... " + name + "!",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
