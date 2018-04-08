package commands

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Set redirects the `set` coommand.
func Set() {

}

// SetWelcomeChannelCommand sets the welcome channel and sends feedback to the user.
func SetWelcomeChannelCommand(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) {

	err := SetWelcomeChannel(db, s, g, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Announce the new welcome channel
	s.ChannelTyping(c.ID)
	_, err = s.ChannelMessageSend(c.ID, "D'accord! <#"+c.ID+"> est maintenant le salon de bienvenue.")
	if err != nil {
		fmt.Println("Couldn't announce the new welcome channel.")
		fmt.Println("Guild : " + g.Name)
		fmt.Println("Channel : " + c.Name)
		fmt.Println(err.Error())
	}
}

// SetWelcomeChannel sets the welcome channel
func SetWelcomeChannel(db *sql.DB, s *discordgo.Session, g *discordgo.Guild, c *discordgo.Channel) error {

	var exists int
	err := db.QueryRow("select count(`welcome`) from `servers` where `server` = ?;", g.ID).Scan(&exists)
	if err != nil {
		fmt.Println("Could not confirm the existence of a welcome channel.")
		fmt.Println("Guild :", g.Name)
		fmt.Println("Channel :", c.Name)
		return err

	} else if exists == 1 {

		// Prepare
		stmt, err := db.Prepare("update `servers` set `welcome` = ? where `server` = ?;")
		if err != nil {
			fmt.Println("Could not prepare to update a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("Channel :", c.Name)
			return err
		}
		defer stmt.Close()

		// Update
		_, err = stmt.Exec(c.ID, g.ID)
		if err != nil {
			fmt.Println("Could not update a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("Channel :", c.Name)
			return err
		}

	} else if exists == 0 {

		// Prepare
		stmt, err := db.Prepare("insert into `servers`(`server`, `welcome`) values(?, ?);")
		if err != nil {
			fmt.Println("Could not prepare to insert a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("Channel :", c.Name)
			return err
		}
		defer stmt.Close()

		// Insert
		_, err = stmt.Exec(g.ID, c.ID)
		if err != nil {
			fmt.Println("Could not insert a welcome channel.")
			fmt.Println("Guild :", g.Name)
			fmt.Println("Channel :", c.Name)
			return err
		}
	}

	return nil
}
