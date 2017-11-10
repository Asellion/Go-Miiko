package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func waitComeBack(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Open channel
	privateChannel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		fmt.Println("Couldn't create a private channel with " + m.User.Username + ".")
		fmt.Println(err.Error())
		return
	}

	// Typing!
	err = s.ChannelTyping(privateChannel.ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get guild
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Couldn't get " + m.User.Username + "'s guild ID.")
		fmt.Println(err.Error())
		return
	}

	// Create an invite's structure
	var invStruct discordgo.Invite
	invStruct.MaxAge = 86400
	invStruct.MaxUses = 1
	invStruct.Temporary = true

	// Look for a valid channel to create an invite
	var invite *discordgo.Invite
	for x := 0; x < len(guild.Channels) && invite == nil; x++ {
		if guild.Channels[x].Type == discordgo.ChannelTypeGuildText {

			// Create invite
			err = nil
			invite, err = s.ChannelInviteCreate(guild.Channels[x].ID, invStruct)
		} else {
			continue
		}
	}
	if err != nil {
		fmt.Println("Couldn't create an invite in " + guild.Name + ".")
		fmt.Println(err.Error())
		return
	}
	if invite == nil {
		fmt.Println("Couldn't create an invite in " + guild.Name + ", but no error message were returned.")
		return
	}

	// Send message
	_, err = s.ChannelMessageSend(privateChannel.ID, "Oh, je suis triste de te voir partir! Si tu veux nous rejoindre à nouveau, j'ai créé une invitation pour toi : https://discord.gg/"+invite.Code)
	if err != nil {
		fmt.Println("Couldn't send the message to " + m.User.Username + "!")
		fmt.Println(err.Error())
	}
}
