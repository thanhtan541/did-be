package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Settings struct {
	Application ApplicationSettings
	//Todo: Secure String
	// HMACSecret secure.String
}

type ApplicationSettings struct {
	Port    uint16
	Host    string
	BaseURL string
	//Todo: Secure String
	// HMACSecret secure.String
}

func LoadConfig() (*Settings, error) {
	v := viper.New()

	// Load base.yaml
	v.SetConfigName("base")
	v.SetConfigType("yaml")
	v.AddConfigPath("configuration")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading base.yaml: %w", err)
	}

	// Merge local.yaml
	v.SetConfigName("local")
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("error merging local.yaml: %w", err)
	}

	var config Settings
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}
