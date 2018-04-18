package config

import (
	"encoding/json"
	"io/ioutil"

	"simple-note-api/domain"
)

func LoadConfig(path string) (domain.Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return domain.Config{}, err
	}

	var config domain.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return domain.Config{}, err
	}

	return config, nil
}
