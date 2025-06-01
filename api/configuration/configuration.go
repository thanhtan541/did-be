package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Settings struct {
	Application ApplicationSettings
	//Todo: Secure String
	// HMACSecret secure.String
}

type ApplicationSettings struct {
	Port uint16
	Host string
	Url  string
	//Todo: Secure String
	// HMACSecret secure.String
}

func LoadConfig() (*Settings, error) {
	dir, err := findGoProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("error getting project root directory: %w", err)
	}

	v := viper.New()

	// Load base.yaml
	v.SetConfigName("base")
	v.SetConfigType("yaml")
	v.AddConfigPath(fmt.Sprintf("%s/%s", dir, "configuration"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading base.yaml: %w", err)
	}

	// Merge local.yaml
	v.SetConfigName("local")
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("error merging local.yaml: %w", err)
	}

	var cfg Settings
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &cfg, nil
}

// In this case, we're looking the parent module
// Note: assumption that configuration live inside
// the target module, "api"
func findGoProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // Found go.mod → this is the project root
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found in any parent directory")
		}

		dir = parent
	}
}
