package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NatoBoram/Go-Miiko/wheel"

	"github.com/bwmarrin/discordgo"
)

func love(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.User) bool {

	// Lover
	lover, err := getLover(s, g)
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

func getLover(s *discordgo.Session, g *discordgo.Guild) (*discordgo.User, error) {

	// User ID
	var userID string
	err := DB.QueryRow("select member from love where server = ?", g.ID).Scan(&userID)
	if err != nil {
		fmt.Println("Couldn't get my lover from this guild.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	// User
	user, err := s.User(userID)
	if err != nil {
		fmt.Println("Couldn't get my dear user.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	// Return
	return user, nil
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
