package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	// Database contains an array of every type of structure out there
	Database database
)

// Contains an array of every type of structure out there
type database struct {
	WelcomeChannels []welcomeChannel
}

// Contains one guild ID with its welcome channel.
type welcomeChannel struct {
	GuildID   string
	ChannelID string
}

// ReadJSON : Reads the JSON database
func ReadJSON() error {
	fmt.Println("Reading the JSON database...")

	// Read a config file
	file, err := ioutil.ReadFile("./database.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Json -> String
	err = json.Unmarshal(file, &Database)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// WriteJSON writes the database to the disk
func WriteJSON() error {
	fmt.Println("Writing the JSON database...")

	// From Database to JSON
	json, err := json.Marshal(Database)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// From JSON to file
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
	return WriteJSON()
}
