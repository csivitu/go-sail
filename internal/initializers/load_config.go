package initializers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/TejasGhatte/go-sail/internal/models"
)

var Config models.Config

func LoadConfig(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		return
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		fmt.Printf("error parsing config file: %v\n", err)
		return
	}
}
