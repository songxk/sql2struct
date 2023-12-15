package main

type dbConfig struct {
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Table       string `yaml:"table"`
	TablePrefix string `yaml:"tablePrefix"`
	ModelSuffix string `yaml:"modelSuffix"`
}

type Config = map[string]*dbConfig
