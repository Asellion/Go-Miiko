package commands

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/NatoBoram/Go-Miiko/wheel"
	"github.com/bwmarrin/discordgo"
)

// Popcorn command
func Popcorn(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check for "pop-corn", "popcorn", "maïs soufflé", "maïs éclaté", "pop corn"
	if strings.Contains(strings.ToLower(m.Content), "popcorn") || strings.Contains(strings.ToLower(m.Content), "pop-corn") || strings.Contains(strings.ToLower(m.Content), "maïs soufflé") || strings.Contains(strings.ToLower(m.Content), "maïs éclaté") || strings.Contains(strings.ToLower(m.Content), "pop corn") {

		// Seed
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		rand := rand.New(source)

		if rand.Float64() <= 1/math.Pow(wheel.Phi(), 1) {

			// It's popcorn time!
			s.ChannelTyping(m.ChannelID)
			_, err := s.ChannelMessageSend(m.ChannelID, getPopcornMessage())
			if err != nil {
				fmt.Println("Couldn't tell everyone how much I love popcorn. Sad :(")
				fmt.Println("Author : " + m.Author.Username)
				fmt.Println("Message : " + m.Content)
				fmt.Println(err.Error())
			}
		}
	}
}

func getPopcornMessage() string {

	// Popcorn Messages
	popcornList := [...]string{

		// Exclamation
		"Popcorn?",
		"Popcorn!",
		"Popcorn?!",
		"Popcorn!?",
		"**Popcorn?!**",
		"**Popcorn!?**",
		"Ah, popcorn!",
		"Hmm, du popcorn...",
		"Hmm, du *popcorn*...",

		// Question
		"On parle de popcorn?",
		"Quelqu'un a dit popcorn?",
		"Quelqu'un a dit popcorn?!",
		"Quelqu'un a dit **popcorn**?!",
		"Quelqu'un a parlé de popcorn?",
		"Quelqu'un a parlé de popcorn?!",
		"Ai-je bien entendu popcorn?",
		"Ai-je bien entendu popcorn?!",
		"Ai-je bien entendu **popcorn**?!",

		// WTF Miiko
		"Moi, j'aime le popcorn!",
		"Le popcorn, c'est génial!",

		// Uhh...
		"Le popcorn, c'est bon et tout, mais il ne faut pas oublier les friandises. J'adore les friandises!",
		"Imagine si... On mélangeait du popcorn... Avec des friandises!",
	}

	// Seed
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)

	// Return
	return popcornList[rand.Intn(len(popcornList))]
}
