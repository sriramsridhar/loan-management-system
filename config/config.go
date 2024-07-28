package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	JWTSecret  string
	GinMode    string
}

var AppConfig *Config

func LoadConfig() {
	if os.Getenv("ENV") == "PRODUCTION" {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // Required if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // Path to look for the config file in

	// Read environment variables
	viper.AutomaticEnv()

	// Set environment variable prefixes and map them to Viper keys
	viper.SetEnvPrefix("app")
	err := viper.BindEnv("DB_USERNAME")
	if err != nil {
		return
	}
	err = viper.BindEnv("DB_PASSWORD")
	if err != nil {
		return
	}
	err = viper.BindEnv("DB_NAME")
	if err != nil {
		return
	}
	err = viper.BindEnv("DB_HOST")
	if err != nil {
		return
	}
	err = viper.BindEnv("DB_PORT")
	if err != nil {
		return
	}
	err = viper.BindEnv("JWT_SECRET")
	if err != nil {
		return
	}
	err = viper.BindEnv("GIN_MODE")
	if err != nil {
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		GinMode:    viper.GetString("GIN_MODE"),
	}
}
