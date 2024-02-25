package adf

import (
	"fmt"
	"io"
	"os"

	coreadf "github.com/gofrontier-com/pony-express/pkg/core/adf"
	"gopkg.in/yaml.v2"
)

func LoadConfig(configFilePath string) (*coreadf.PonyConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := coreadf.PonyConfig{}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &cfg, nil
}
