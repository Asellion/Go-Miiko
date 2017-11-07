package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Public variables
	Token     string
	BotPrefix string

	// Private variables
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
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

	// It works!
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
