package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type IChainConfig interface {
	GetId() string      //String identifier for the chain for example ethereum-sepolia
	GetChainId() uint64 //Integer identifier for the chain for example 11155111
	GetFamily() string  //Family of the chain for example evm
	GetName() string    //Name of the chain for example Ethereum Sepolia
}
type ChainFamily map[uint64]IChainConfig
type Config struct {
	ConfigPath string `mapstructure:"CONFIG_PATH"`
	//ActiveChains      map[string]bool        `mapstructure:"active_chains"` //Store all active chains in the scalar network
}

var GlobalConfig Config

// func loadConfig() (Config, error) {
// 	var config Config

// 	// Automatically map ENV variables
// 	viper.AutomaticEnv()

// 	// Optionally set default values
// 	viper.SetDefault("APP_NAME", "MyApp")
// 	viper.SetDefault("PORT", 8080)
// 	viper.SetDefault("DEBUG", false)

// 	// Unmarshal into the config struct
// 	err := viper.Unmarshal(&config)
// 	return config, err
// }

func LoadEnv(environment string) error {
	// Tell Viper to read from environment
	viper.AutomaticEnv()
	viper.Set("CONFIG_PATH", os.Getenv("CONFIG_PATH"))
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Add support for .env files
	// if environment == "" {
	// 	viper.SetConfigFile(".env") // Set config file
	// } else {
	// 	viper.SetConfigName(environment)
	// 	viper.SetConfigType("env")
	// }
	// viper.AddConfigPath(".") // look for config in the working directory
	// Read the .env file
	// if err := viper.ReadInConfig(); err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		// Config file not found; ignore error if desired
	// 		fmt.Printf("No %s.env file found", environment)
	// 	} else {
	// 		// Config file was found but another error was produced
	// 		return fmt.Errorf("error reading config file: %w", err)
	// 	}
	// }
	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}
	log.Info().Msgf("Loaded config: %+v", GlobalConfig)
	return nil
}

func ReadJsonArrayConfig[T any](filePath string) ([]T, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Unmarshal directly into slice
	result, err := ParseJsonArrayConfig[T](content)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config from %s: %w", filePath, err)
	}
	return result, nil
}

func ReadJsonConfig[T any](filePath string) (*T, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Unmarshal directly into slice
	result, err := ParseJsonConfig[T](content)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config from %s: %w", filePath, err)
	}

	return result, nil
}

func ParseJsonArrayConfig[T any](content []byte) ([]T, error) {
	var result []T
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func ParseJsonConfig[T any](content []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
