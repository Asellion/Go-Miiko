package bot

import (
	"fmt"
	"math"

	"github.com/bwmarrin/discordgo"
)

func pin(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Get channel structure
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure of a MessageReactionAdd!")
		fmt.Println("m.ChannelID : " + m.ChannelID)
		fmt.Println(err.Error())
		return
	}

	// DM?
	if channel.Type == discordgo.ChannelTypeDM {
		return
	}

	// Get the message structure
	message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		fmt.Println("Couldn't get the message structure of a MessageReactionAdd!")
		fmt.Println("m.ChannelID : " + m.ChannelID)
		fmt.Println("m.MessageID : " + m.MessageID)
		fmt.Println(err.Error())
		return
	}

	// Get the guild structure
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Couldn't get the guild structure of a MessageReactionAdd!")
		fmt.Println("channel.GuildID : " + channel.GuildID)
		fmt.Println(err.Error())
		return
	}

	// Get phi
	// phi := (1 + math.Sqrt(5)) / 2

	// Get people online
	var onlineCount int
	var notOfflineCount int
	for x := 0; x < len(guild.Presences); x++ {
		if guild.Presences[x].Status == discordgo.StatusOnline {
			onlineCount++
		}
		if guild.Presences[x].Status != discordgo.StatusOffline {
			notOfflineCount++
		}
	}

	// Get the reactions
	var singleReactionCount int
	var totalReactionsCount int
	for x := 0; x < len(message.Reactions); x++ {
		singleReactionCount = int(math.Max(float64(singleReactionCount), float64(message.Reactions[x].Count)))
		totalReactionsCount += message.Reactions[x].Count
	}

	// Pins needs at least 3 reactions!
	var absoluteMinimum float64 = 3

	// Get minimum for pin
	minOnline := int(math.Max(absoluteMinimum, math.Ceil(math.Sqrt(float64(onlineCount)))))
	// minTotal := int(math.Max(minOnline + 1, math.Ceil(math.Sqrt(float64(notOfflineCount)))))

	// Count the reactions
	if singleReactionCount >= minOnline {
		// || totalReactionsCount >= minTotal

		// Pin it!
		err := s.ChannelMessagePin(m.ChannelID, m.MessageID)
		if err != nil {
			fmt.Println("Couldn't pin a popular message!")
			fmt.Println("Message : " + message.Content)
			fmt.Println(err.Error())
			return
		}
	}
}
