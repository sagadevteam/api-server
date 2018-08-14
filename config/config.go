package config

import (
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

var (

	// API - Api general config
	API APIConf

	// DB - DB general config
	DB Database
)

// Database is the struct of cassandra config unmarshal
type Database struct {
	Table string `yaml:"db_table"`
	Host  string `yaml:"db_host"`
	User  string `yaml:"db_user"`
	Pwd   string `yaml:"db_pwd"`
	Port  int    `yaml:"db_port"`
}

// APIConf config
type APIConf struct {
	APIprotocol string `yaml:"api_protocol"`
	Domain      string `yaml:"api_domain"`
	Port        string `yaml:"api_port"`
}

// DeploySet - deploy config
type DeploySet struct {
	Database Database `yaml:"db"`
	API      APIConf  `yaml:"api"`
}

// ReadConf read the config file from input filePath
func init() {
	var dconf DeploySet
	if content, ioErr := ioutil.ReadFile("./config/saga.conf.yaml"); ioErr != nil {
		logrus.Fatalf("read service config file error: %v", ioErr)
	} else {
		if ymlErr := yaml.Unmarshal(content, &dconf); ymlErr != nil {
			log.Fatalf("error while unmarshal from db config: %v", ymlErr)
		}
	}

	API = dconf.API
	DB = dconf.Database

}
