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

	var (
		userID string
		pins   int
	)

	// Select potential lovers
	rows, err := DB.Query("select `member`, `count` from `love` where `server` = ? order by `count` desc;", g.ID)
	if err != nil {
		fmt.Println("Couldn't get my lovers from this guild.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}
	defer rows.Close()

	// For each rows
	for rows.Next() {

		// Scan it
		err := rows.Scan(&userID, &pins)
		if err != nil {
			fmt.Println("Couldn't scan a potential lover.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// User
		user, err := s.User(userID)
		if err != nil {
			fmt.Println("Couldn't get a potential lover's user.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// Member
		member, err := s.GuildMember(g.ID, user.ID)
		if err != nil {
			fmt.Println("Couldn't get a potential lover's member.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// Owner
		if g.OwnerID == user.ID {
			continue
		}

		// Roles
		if len(member.Roles) == 1 {
			return user, nil
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Couldn't loop my lovers.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	// Unreachable code.
	user, err := s.User(g.OwnerID)
	if err != nil {
		fmt.Println("Couldn't love the owner.")
		fmt.Println("Guild :", g.Name)
		return nil, err
	}

	return user, err
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
