package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func loadConfig(filePath string) (Config, error) {
	configFile := ""
	if filePath != "" {
		configFile = filePath
	} else {
		execPath, err := GetExecPath()
		if err != nil {
			log.Fatal(err.Error())
		}
		configFile = execPath + "/sql2struct.yaml"
	}
	if _, err := os.Stat(configFile); err != nil && errors.Is(err, os.ErrNotExist) {
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
