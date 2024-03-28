package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		// 서버 실행 모드
		RunMode string `default:"dev" yaml:"runMode"`
		// 서버 실행 포트
		Port int `default:"5000" yaml:"port"`
	} `yaml:"server"`
	Database struct {
	} `yaml:"database"`
}

func (conf *Config) GetConfig() *Config {
	ymlFile, err := os.ReadFile("config.yml")

	if err != nil {
		log.Fatal("Can't find config.yml")
	}

	err = yaml.Unmarshal(ymlFile, conf)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// if strings.ToUpper(conf.Server.RunMode) == "DEV" {
	// 	fmt.Println(conf)
	// }
	return conf
}
