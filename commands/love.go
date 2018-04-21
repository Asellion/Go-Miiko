package commands

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// GetLoverCmd outputs the lover
func GetLoverCmd(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, u *discordgo.User) {

	// Owner only!
	if u.ID != g.OwnerID {
		return
	}

	// Inform the user that I'm typing
	s.ChannelTyping(c.ID)

	// Get lover
	lover, err := GetLover(db, s, g)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Send response
	_, err = s.ChannelMessageSend(c.ID, getLoverMessage(lover))
	if err != nil {
		fmt.Println("Couldn't reveal my lover.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		fmt.Println("User :", u.Username)
	}
}

// GetLover gets this guild's lover.
func GetLover(db *sql.DB, s *discordgo.Session, g *discordgo.Guild) (*discordgo.User, error) {

	var (
		userID string
		pins   int
	)

	// Select potential lovers
	rows, err := db.Query("select `member`, `count` from `pins-count` where `server` = ? order by `count` desc;", g.ID)
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

		// Member
		member, err := s.GuildMember(g.ID, userID)
		if err != nil {
			fmt.Println("Couldn't get a potential lover's member.")
			fmt.Println("Guild :", g.Name)
			continue
		}

		// Owner
		if g.OwnerID == member.User.ID {
			continue
		}

		// Roles
		if len(member.Roles) == 1 {
			return member.User, nil
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

func getLoverMessage(u *discordgo.User) string {

	// Messages
	loveList := [...]string{
		"Disons que je chéris particulièrement <@" + u.ID + ">.",
		"Si j'avais à marier quelqu'un... Ce serait <@" + u.ID + ">!",
		"Peut-être... <@" + u.ID + ">?",
	}

	// Seed
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)

	// Return
	return loveList[seed.Intn(len(loveList))]
}
