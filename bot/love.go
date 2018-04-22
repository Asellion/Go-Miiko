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
		":heart:",
		"\\*Frissonne*",
		"Coucou <@" + u.ID + "> :3",
		"<@" + u.ID + ">-senpai \\*-*",
		"\\*-*",
		"<@" + u.ID + ">... Tu veux qu'on fasse quelque chose ensemble?",
		"Oh, <@" + u.ID + ">, est-ce que je te manque?",
		"Est-ce que tu penses à moi, <@" + u.ID + ">?",
		"Tiens-moi la main, <@" + u.ID + ">",
		"J'ai trouvé un morceau de cristal pour toi, <@" + u.ID + "> :heart:",
		"Cette voix est une musique à mes oreilles",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
