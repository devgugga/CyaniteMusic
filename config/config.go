package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Discord struct {
		Token   string `yaml:"token"`
		GuildID string `yaml:"guild_id"`
	} `yaml:"discord"`

	PocketBase struct {
		URL    string `yaml:"url"`
		APIKey string `yaml:"api_key"`
	} `yaml:"pocket-base"`

	YouTube struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"youtube"`
}

var Cfg *Config

// LoadConfig reads and parses the configuration file located at "config/config.yaml".
// The configuration file is expected to be in YAML format and contains the necessary
// credentials and settings for the Discord, PocketBase, and YouTube integrations.
//
// The function opens the configuration file, decodes its contents into a Config struct,
// and assigns the pointer to the global variable Cfg. If any errors occur during file
// opening or decoding, the function logs the error and terminates the program.
func LoadConfig() {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&Cfg); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}
}
