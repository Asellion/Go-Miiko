package main

import (
	"database/sql"
	"fmt"

	"github.com/NatoBoram/Go-Miiko/bot"
	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database config.DBStruct
	discord  config.DiscordStruct
	db       *sql.DB // http://go-database-sql.org/
	session  *discordgo.Session
)

func main() {

	// License
	fmt.Println("")
	fmt.Println("Go-Miiko : Manages an Eldarya-themed Discord server.")
	fmt.Println("Copyright © 2018 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println("Contact : https://github.com/NatoBoram/Go-Miiko")
	fmt.Println("")

	// Read the database config
	err := config.ReadDB(&database)
	if err != nil {
		fmt.Println("Could not load the database configuration.")
		fmt.Println(err.Error())
		config.WriteTemplateDB()
		return
	}

	// Read the Discord config
	err = config.ReadDiscord(&discord)
	if err != nil {
		fmt.Println("Could not load the Discord configuration.")
		fmt.Println(err.Error())
		config.WriteTemplateDiscord()
		return
	}

	// Connect to MariaDB
	db, err = sql.Open("mysql", database.User+":"+database.Password+"@tcp("+database.Address+":"+database.Port+")/"+database.Database)
	if err != nil {
		fmt.Println("Could not connect to the database.")
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Create a Discord session
	session, err = discordgo.New("Bot " + discord.Token)
	if err != nil {
		fmt.Println("Could not create a Discord session.")
		fmt.Println(err.Error())
		return
	}

	// Connect to Discord
	err = session.Open()
	if err != nil {
		fmt.Println("Could not connect to Discord.")
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	// Give this bot some life!
	err = bot.Start(db, session, discord.MasterID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Wait for future input
	<-make(chan struct{})
	return
}
