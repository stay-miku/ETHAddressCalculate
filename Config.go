package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Thread  int      `yaml:"thread"`
	RunTime int      `yaml:"run_times"`
	Reg     []string `yaml:"reg"`
	Output  string   `yaml:"output"`
}

var Config config

func init() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
