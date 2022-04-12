package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Mysql struct {
		Username string `yaml:"username"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Param    string `yaml:"param"`
		Dbname   string `yaml:"dbname"`
	}
}

func Initial(con *Config) {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, con)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
