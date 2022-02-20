package main

import (
    "fmt"
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func log_error(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

type Config struct {
	Servers map[string]Server
}

type Server struct {
	IP string `yaml:"ip"`
	PORT string `yaml:"port"`
}

func main() {

	yml_file, err := ioutil.ReadFile("config.yml")

	log_error(err)

	config := Config{}

	err := yaml.Unmarshal(yml_file, &config)

	log_error(err)

	for ip, port := range config {
		fmt.Printf("%s: %s\n", ip, port)
	}

}
