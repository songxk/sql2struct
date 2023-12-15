package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func loadConfig(filePath string) (Config, error) {
	configFile := GetRootPath() + "/config.yaml"
	if filePath != "" {
		configFile = filePath
	}
	if _, err := os.Stat(filePath); err != nil && errors.Is(err, os.ErrNotExist) {
		log.Fatal(configFile + " not exist ")
	}
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Reading config file failed, " + err.Error())
	}
	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal("Unmarshal config file failed, " + err.Error())
	}
	return c, nil
}
