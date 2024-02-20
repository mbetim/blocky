package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mbetim/blocky/cmd"
	"github.com/spf13/viper"
)

func main() {
	StartConfig()
	cmd.Execute()
}

func StartConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configBasePath := filepath.Join(homeDir, ".config", "blocky")
	configFullPath := filepath.Join(configBasePath, "config.yaml")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configBasePath)

	if _, err := os.Stat(configBasePath); os.IsNotExist(err) {
		if err := os.Mkdir(configBasePath, os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(configFullPath); os.IsNotExist(err) {
		if _, err := os.Create(configFullPath); err != nil {
			return err
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, ", err)
		return err
	}

	return nil
}
