package main

import (
	"database/sql"

	"github.com/NatoBoram/Go-Miiko/bot"
	"github.com/NatoBoram/Go-Miiko/config"
	"github.com/bwmarrin/discordgo"
)

var (
	database config.DBStruct
	discord  config.DiscordStruct
	db       *sql.DB
	session  *discordgo.Session
)

func main() {

	// Read the database config
	err := config.ReadDB(&database)
	if err != nil {
		println("Could not load the database configuration.")
		println(err.Error())
		config.WriteTemplateDB()
		return
	}

	// Read the Discord config
	err = config.ReadDiscord(&discord)
	if err != nil {
		println("Could not load the Discord configuration.")
		println(err.Error())
		config.WriteTemplateDiscord()
		return
	}

	// Connect to MariaDB
	db, err = sql.Open("mysql", database.User+":"+database.Password+"@tcp("+database.Address+":"+database.Port+")/"+database.Database)
	if err != nil {
		println("Could not connect to the database.")
		println(err.Error())
		return
	}
	defer db.Close()

	// Create a Discord session
	session, err = discordgo.New("Bot " + discord.Token)
	if err != nil {
		println("Could not create a Discord session.")
		println(err.Error())
		return
	}

	// Connect to Discord
	err = session.Open()
	if err != nil {
		println("Could not connect to Discord.")
		println(err.Error())
		return
	}

	// License
	println("")
	println("Go-Miiko : Manages an Eldarya-themed Discord server.")
	println("Copyright © 2018 Nato Boram")
	println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	println("Contact : https://github.com/NatoBoram/Go-Miiko")
	println("")

	// Give this bot some life!
	bot.Start(db, session)

	// Wait for future input
	<-make(chan struct{})
	return
}
