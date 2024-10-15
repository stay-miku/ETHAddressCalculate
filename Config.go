package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	GoThread         int      `yaml:"go-thread"`
	ETHPhraseThread  int      `yaml:"eth-phrase-thread"`
	ETHKeyThread     int      `yaml:"eth-key-thread"`
	TRONPhraseThread int      `yaml:"tron-phrase-thread"`
	TRONKeyThread    int      `yaml:"tron-key-thread"`
	RunTime          int      `yaml:"run-times"`
	ETHReg           []string `yaml:"eth-reg"`
	TronReg          []string `yaml:"tron-reg"`
	ETHSuffix        int      `yaml:"eth-suffix"`
	TronSuffix       int      `yaml:"tron-suffix"`
	ETHPrefix        int      `yaml:"eth-prefix"`
	TronPrefix       int      `yaml:"tron-prefix"`
	ETHPS            []int    `yaml:"eth-p-s"`
	TronPS           []int    `yaml:"tron-p-s"`
	Output           string   `yaml:"output"`
	Length           int      `yaml:"len"`
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
