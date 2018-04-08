package commands

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Get redirects the `get` coommand.
func Get() {

}

// GetWelcomeChannelCommand send the welcome channel to an user.
func GetWelcomeChannelCommand(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	welcome, err := GetWelcomeChannel(db, s, g, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Send the welcome channel
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "Le salon de bienvenue est <#"+welcome.ID+">.")
	if err != nil {
		fmt.Println("Couldn't send a message.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
		return
	}
}

// GetWelcomeChannel gets the welcome channel
func GetWelcomeChannel(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) (*discordgo.Channel, error) {

	// Does it exists?
	var exists int
	err := db.QueryRow("select count(`welcome`) from `servers` where `server` = ?;", g.ID).Scan(&exists)
	if err != nil {
		fmt.Println("Could not confirm the existence of a welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		return nil, err

	} else if exists == 0 {

		// Set this one if it doesn't exist.
		err := SetWelcomeChannel(db, s, g, c)
		if err != nil {
			return nil, err
		}
		return c, nil

	} else if exists == 1 {

		// Get the welcome channel's ID
		var welcome string
		err = db.QueryRow("select `welcome` from `servers` where `server` = ?;", g.ID).Scan(&welcome)
		if err != nil {
			fmt.Println("Could not select a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("Channel :", c.Name)
			return nil, err
		}

		// Does the channel still exists?
		channel, err := s.Channel(welcome)
		if err != nil {
			fmt.Println("Couldn't get the channel structure of a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("ChannelID : " + welcome)
			fmt.Println(err.Error())

			// Set this one if it doesn't exist.
			err := SetWelcomeChannel(db, s, g, c)
			if err != nil {
				return nil, err
			}
			return c, nil
		}

		// It exists!
		return channel, nil
	}

	// Unreachable code.
	return c, err
}
