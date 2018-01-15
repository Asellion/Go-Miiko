package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	// Database contains an array of every type of structure out there
	Database database
)

// Contains an array of every type of structure out there
type database struct {
	Token           string
	MasterID        string
	WelcomeChannels map[string]string
}

// ReadJSON : Reads the JSON database
func ReadJSON() error {

	// Read the JSON database
	file, err := ioutil.ReadFile("./database.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Json -> Database
	err = json.Unmarshal(file, &Database)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// WriteJSON : Writes the database to the disk
func WriteJSON() error {

	// From Database to JSON
	json, err := json.Marshal(Database)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// From JSON to File
	err = ioutil.WriteFile("./database.json", json, os.FileMode(int(0777)))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// WriteNewJSON writes the database to the disk
func WriteNewJSON() error {
	fmt.Println("Creating a new database.")
	var newDatabase database
	Database = newDatabase
	Database.WelcomeChannels = make(map[string]string)
	return WriteJSON()
}

// UpdateWelcomeChannel gets an automatic welcome message and updates the welcome channel with this.
func UpdateWelcomeChannel(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure!")
		fmt.Println(err.Error())
		return
	}

	// Create a new map
	if len(Database.WelcomeChannels) == 0 {
		Database.WelcomeChannels = make(map[string]string)
	}

	// Update the value
	Database.WelcomeChannels[channel.GuildID] = channel.ID

	// Save
	WriteJSON()
}
