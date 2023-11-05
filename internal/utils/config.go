package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(fileName string) (*Config, error) {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Printf("Failed to read YAML file: %v", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Printf("Failed to unmarshal YAML: %v", err)
		return nil, err
	}
	return config, err
}
