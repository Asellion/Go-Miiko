package bot

import (
	"fmt"
	"math"

	"github.com/NatoBoram/Go-Miiko/wheel"
	"github.com/bwmarrin/discordgo"
)

func pin(s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel, m *discordgo.Message) {

	// DM?
	if c.Type == discordgo.ChannelTypeDM {
		return
	}

	// Get people online
	var onlineCount int
	for x := 0; x < len(g.Presences); x++ {
		if g.Presences[x].Status == discordgo.StatusOnline {
			onlineCount++
		}
	}

	// Get the reactions
	var singleReactionCount int
	for x := 0; x < len(m.Reactions); x++ {
		singleReactionCount = wheel.MinInt(singleReactionCount, m.Reactions[x].Count)
	}

	// Pins needs at least 3 reactions!
	var absoluteMinimum float64 = 3

	// Get minimum for pin
	minOnline := int(math.Max(absoluteMinimum, math.Ceil(math.Sqrt(float64(onlineCount)))))

	// Check the reactions
	if singleReactionCount >= minOnline {

		// Pin it!
		err := s.ChannelMessagePin(c.ID, m.ID)
		if err != nil {
			fmt.Println("Couldn't pin a popular message!")
			fmt.Println("Guild : " + g.Name)
			fmt.Println("Channel : " + c.Name)
			fmt.Println("Author : " + m.Author.Username)
			fmt.Println("Message : " + m.Content)
			fmt.Println(err.Error())
			return
		}

		// Add it to database!
		pindb(g, m)
	}
}

// Add a single pin to the database.
func pindb(g *discordgo.Guild, m *discordgo.Message) {

	// Prepare
	stmt, err := DB.Prepare("insert into `pins`(`channel`, `member`, `message`) values(?, ?, ?)")
	if err != nil {
		fmt.Println("Couldn't prepare a pin.")
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	// Execute
	_, err = stmt.Exec(g.ID, m.Author.ID, m.ID)
	if err != nil {
		fmt.Println("Couldn't insert a pin.")
		fmt.Println(err.Error())
		return
	}
}
