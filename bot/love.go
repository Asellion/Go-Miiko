package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NatoBoram/Go-Miiko/commands"
	"github.com/NatoBoram/Go-Miiko/wheel"

	"github.com/bwmarrin/discordgo"
)

func love(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.User) bool {

	// Lover
	lover, err := commands.GetLover(DB, s, g)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Verify if it's the one true love
	if u.ID == lover.ID {

		// Random
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		rand := rand.New(source)
		random := rand.Float64()

		// Rate Limit
		if random < 1/(wheel.Phi()*10) {

			// Give some love!
			s.ChannelTyping(c.ID)
			_, err = s.ChannelMessageSend(c.ID, getLoveMessage(u))
			if err != nil {
				fmt.Println("Couldn't express my love.")
				fmt.Println("Guild :")
				fmt.Println("Channel :")
				fmt.Println("User :")
				fmt.Println("Message :")
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
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
