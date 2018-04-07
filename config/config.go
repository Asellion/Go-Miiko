package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DBStruct hosts the bot's database configuration.
type DBStruct struct {
	User     string
	Password string
	Address  string
	Port     string
	Database string
}

// DiscordStruct hosts the bot's Discord configuration.
type DiscordStruct struct {
	Token    string
	MasterID string
}

// ReadDB reads the database configuration.
func ReadDB(dbConfig *DBStruct) error {

	// Read the JSON file
	file, err := ioutil.ReadFile("./Miiko/db.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &dbConfig)
	if err != nil {
		return err
	}

	return nil
}

// ReadDiscord reads the Discord configuration.
func ReadDiscord(discordConfig *DiscordStruct) error {

	// Read the JSON file
	file, err := ioutil.ReadFile("./Miiko/discord.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &discordConfig)
	if err != nil {
		return err
	}

	return nil
}

// WriteDB : Writes the database configuration to the disk.
func WriteDB(dbConfig *DBStruct) error {

	// From Database to JSON
	json, err := json.Marshal(dbConfig)
	if err != nil {
		return err
	}

	// From JSON to File
	err = ioutil.WriteFile("./Miiko/db.json", json, os.FileMode(int(0777)))
	if err != nil {
		return err
	}

	return nil
}

// WriteDiscord : Writes the Discord configuration to the disk.
func WriteDiscord(discordConfig *DiscordStruct) error {

	// From Database to JSON
	json, err := json.Marshal(discordConfig)
	if err != nil {
		return err
	}

	// From JSON to File
	err = ioutil.WriteFile("./Miiko/db.json", json, os.FileMode(int(0777)))
	if err != nil {
		return err
	}

	return nil
}

// WriteTemplateDB creates a template db confituration with null values and writes them to the disk to allow the user to insert its own values.
func WriteTemplateDB() error {
	var newDB DBStruct
	return WriteDB(&newDB)
}

// WriteTemplateDiscord creates a template Discord confituration with null values and writes them to the disk to allow the user to insert its own values.
func WriteTemplateDiscord() error {
	var newDiscord DiscordStruct
	return WriteDiscord(&newDiscord)
}
