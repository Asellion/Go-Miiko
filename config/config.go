package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Public variables
	Token              string
	BotPrefix          string
	BotMasterChannelID string

	// Private variables
	config *configStruct
)

type configStruct struct {
	Token              string `json:"Token"`
	BotPrefix          string `json:"BotPrefix"`
	BotMasterChannelID string `json:"BotMasterChannelID"`
}

// ReadConfig : Reads the config of config.json.
func ReadConfig() error {
	fmt.Println("Reading config file...")

	// Read a config file
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Json -> String
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Export
	Token = config.Token
	BotPrefix = config.BotPrefix
	BotMasterChannelID = config.BotMasterChannelID
	return nil
}
