package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppPort      string        `yaml:"port"`
	AppHost      string        `yaml:"host"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	DbConnString string        `yaml:"db_conn_string"`
}

func Parse() (*Config, error) {
	filename, err := filepath.Abs("./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("can't get config path: %s", err.Error())
	}
	yamlConf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read conf: %s", err.Error())
	}

	var config Config
	err = yaml.Unmarshal(yamlConf, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall conf: %s", err.Error())
	}

	return &config, nil
}
