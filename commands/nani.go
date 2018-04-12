package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Nani command
func Nani(s *discordgo.Session, m *discordgo.Message) bool {

	// Check for "Omae Wa Mou Shindeiru"
	if strings.Contains(strings.ToLower(m.Content), "omae wa mou shindeiru") {

		// Nani?!
		_, err := s.ChannelMessageSend(m.ChannelID, getNaniMessage())
		if err != nil {
			fmt.Println("Couldn't express my surprise. Sad :(")
			fmt.Println(err.Error())
			return false
		}
		return true
	}
	return false
}

func getNaniMessage() string {

	// Nani Messages
	var naniList []string

	naniList = append(naniList, "Nani?")
	naniList = append(naniList, "Nani?!")
	naniList = append(naniList, "Nani!?")
	naniList = append(naniList, "**Nani?!**")
	naniList = append(naniList, "**Nani!?**")
	naniList = append(naniList, "**Nani ?!?**")
	naniList = append(naniList, "**Nani !?!**")

	// Seed
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return naniList[rand.Intn(len(naniList))]
}
