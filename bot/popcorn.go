package bot

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func popcorn(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check for "pop-corn", "popcorn", "maïs soufflé", "maïs éclaté", "pop corn"
	if strings.Contains(strings.ToLower(m.Content), "popcorn") || strings.Contains(strings.ToLower(m.Content), "pop-corn") || strings.Contains(strings.ToLower(m.Content), "maïs soufflé") || strings.Contains(strings.ToLower(m.Content), "maïs éclaté") || strings.Contains(strings.ToLower(m.Content), "pop corn") {

		// Get phi
		phi := (1 + math.Sqrt(5)) / 2

		// Seed
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		rand := rand.New(source)

		if rand.Float64() > 1/math.Pow(phi, 2) {

			// Typing!
			err := s.ChannelTyping(m.ChannelID)
			if err != nil {
				fmt.Println(err.Error())
			}

			// It's popcorn time!
			_, err = s.ChannelMessageSend(m.ChannelID, getPopcornMessage())
			if err != nil {
				fmt.Println("Couldn't tell everyone how much I love popcorn. Sad :(")
				fmt.Println(err.Error())
			}
		}
	}
}

func getPopcornMessage() string {

	// Popcorn Messages
	var popcornList []string

	// Exclamation
	popcornList = append(popcornList, "Popcorn?")
	popcornList = append(popcornList, "Popcorn!")
	popcornList = append(popcornList, "Popcorn?!")
	popcornList = append(popcornList, "Popcorn!?")

	// Surprize
	popcornList = append(popcornList, "Quelqu'un a dit popcorn?")
	popcornList = append(popcornList, "Quelqu'un a dit popcorn?!")
	popcornList = append(popcornList, "Quelqu'un a dit **popcorn**?!")
	popcornList = append(popcornList, "Ai-je bien entendu popcorn?")
	popcornList = append(popcornList, "Ai-je bien entendu popcorn?!")
	popcornList = append(popcornList, "Ai-je bien entendu **popcorn**?!")

	// Question
	popcornList = append(popcornList, "On parle de popcorn?")
	popcornList = append(popcornList, "Quelqu'un a parlé de popcorn?")
	popcornList = append(popcornList, "Quelqu'un a parlé de popcorn?!")

	// WTF Miiko
	popcornList = append(popcornList, "Moi, j'aime le popcorn!")

	// Seed
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return popcornList[rand.Intn(len(popcornList))]
}
