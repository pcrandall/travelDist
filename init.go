package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pcrandall/travelDist/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SheetName string `yaml:"sheetname"`
	Levels    []struct {
		Floor   int `yaml:"floor"`
		Navette []struct {
			Name string `yaml:"name"`
			IP   string `yaml:"ip"`
			Row  string `yaml:"row"`
		} `yaml:"navette"`
	} `yaml:"levels"`
}

func init() {
	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}

	dbpath := filepath.Join(".", "db")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		os.MkdirAll(dbpath, os.ModePerm)
	}
	GetConfig()
}

func GetConfig() {
	if _, err := os.Stat("./config/config.yml"); err == nil { // check if config file exists
		yamlFile, err := ioutil.ReadFile("./config/config.yml")
		utils.CheckErr(err, "init: error reading ./config/config.yml")
		err = yaml.Unmarshal(yamlFile, &config)
		utils.CheckErr(err, "init: error unmarshalling yaml into &config")
	} else if os.IsNotExist(err) { // config file not included, use embedded config
		yamlFile, err := Asset("./config/config.yml")
		utils.CheckErr(err, "init: error reading Asset(./config/config.yml)")
		err = yaml.Unmarshal(yamlFile, &config)
		utils.CheckErr(err, "init: error unmarshalling yaml into &config")
	} else {
		log.Panic("Schrodinger: file may or may not exist. See err for details.")
	}
}
