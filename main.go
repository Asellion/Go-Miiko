package main

import (
	"fmt"

	"github.com/NatoBoram/Go-Miiko/bot"
	"github.com/NatoBoram/Go-Miiko/config"
)

func main() {

	// Reads the configuration
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Reads the JSON database
	err = config.ReadJSON()
	if err != nil {
		fmt.Println(err.Error())
		config.WriteNewJSON()
		// return
	}

	// License
	fmt.Println("")
	fmt.Println("Go-Miiko : Manages an Eldarya-themed Discord server.")
	fmt.Println("Copyright © 2017 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println("Contact : https://github.com/NatoBoram/Go-Miiko")
	fmt.Println("")

	// Give this bot some life!
	bot.Start()

	// Wait for future input
	<-make(chan struct{})
	return
}
