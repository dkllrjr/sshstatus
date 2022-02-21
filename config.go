package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Server struct {
	HOST string `yaml:"host"`
	PORT string `yaml:"port"`
}

func GetConfig(config_file string) map[string]Server {

	config_bytes, err := ioutil.ReadFile(config_file)
	LogError(err)

	config := make(map[string]Server)

	err = yaml.Unmarshal(config_bytes, &config)
	LogError(err)

	return config
}
