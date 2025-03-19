package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Port     int    `mapstructure:"port"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	possiblePaths := []string{
		currentDir,
		filepath.Join(currentDir, "configs"),
		filepath.Join(currentDir, "..", "configs"),
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")

	for _, path := range possiblePaths {
		viper.AddConfigPath(path)
		log.Println("Checking config path:", path)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	log.Println("Configuration loaded successfully.")
	return &config, nil

}
