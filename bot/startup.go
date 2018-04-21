package bot

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func refresh(db *sql.DB, s *discordgo.Session) {

	// Wait for a minute
	time.Sleep(time.Minute)

	// Begin
	start := time.Now()
	fmt.Println("Begin refresh :", start.String())

	for _, guild := range s.State.Guilds {
		for _, channel := range guild.Channels {

			// Ignore non-text channels
			if channel.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Pins
			pins, err := s.ChannelMessagesPinned(channel.ID)
			if err != nil {
				fmt.Println("Couldn't get a channel's pins.")
				fmt.Println("Guild :", guild.Name)
				fmt.Println("Channel :", channel.Name)
				fmt.Println(err.Error())
				continue
			}

			// Pin
			for _, message := range pins {
				pindb(guild, message)
			}
		}
	}

	// End
	duration := time.Since(start)
	fmt.Println("Finished refresh :", duration.String())
}
