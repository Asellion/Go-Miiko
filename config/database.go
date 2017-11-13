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

// UpdateWelcomeChannel gets an automatic welcome message and updates the welcome channel with this.
func UpdateWelcomeChannel(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get channel structure
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Couldn't get the channel structure!")
		fmt.Println(err.Error())
		return
	}

	// Is there already an entry for this Guild?
	index := len(Database.WelcomeChannels)
	for x := 0; x < len(Database.WelcomeChannels); x++ {
		if Database.WelcomeChannels[x].GuildID == channel.GuildID {
			index = x
			break
		}
	}

	// Nope
	if index == len(Database.WelcomeChannels) {

		// Create entry
		var newWelcomeChannel welcomeChannel
		newWelcomeChannel.GuildID = channel.GuildID
		newWelcomeChannel.ChannelID = channel.ID
		Database.WelcomeChannels = append(Database.WelcomeChannels, newWelcomeChannel)

	} else {

		// Update entry
		Database.WelcomeChannels[index].ChannelID = channel.ID
	}

	// Save
	WriteJSON()
}

// GetWelcomeChannelByGuildID outputs a guild's welcome channel ID. Watch out for "" value!
func GetWelcomeChannelByGuildID(guildID string) string {

	// Is there already an entry for this Guild?
	index := len(Database.WelcomeChannels)
	for x := 0; x < len(Database.WelcomeChannels); x++ {
		if Database.WelcomeChannels[x].GuildID == guildID {
			index = x
			break
		}
	}

	// Nope
	if index == len(Database.WelcomeChannels) {
		return ""
	}

	// Yes!
	return Database.WelcomeChannels[index].ChannelID
}
