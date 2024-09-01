package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Discord struct {
		Token   string `yaml:"token"`
		GuildID string `yaml:"guild-id"`
	}
	Lavalink struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	}
}

var Cfg *Config

func LoadConfig() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Cfg); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}
}
