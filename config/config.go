package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (

	// Token : Token.
	Token string

	// BotPrefix : Command used to call the bot. Unused.
	BotPrefix string

	// BotMasterChannelID : ChannelID of the BotMaster. I have to change it for MasterID... Eventually.
	BotMasterChannelID string

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
