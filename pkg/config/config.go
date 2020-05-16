package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ServerStruct for server fields
type ServerStruct struct {
	ListenAddr string `yaml:"listen_addr"`
}

// DBStruct for db fields
type DBStruct struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Name   string `yaml:"name"`
}

// LogStruct for log fields
type LogStruct struct {
	OutputLevel        string `yaml:"output_level"`
	OutputPath         string `yaml:"output_path"`
	RotationPath       string `yaml:"rotation_path"`
	RotationMaxSize    int    `yaml:"rotation_max_size"`
	RotationMaxAge     int    `yaml:"rotation_max_age"`
	RotationMaxBackups int    `yaml:"rotation_max_backups"`
	JSONEncoding       bool   `yaml:"json_encoding"`
}

// Config structure for server
type Config struct {
	Server ServerStruct `yaml:"server"`
	DB     DBStruct     `yaml:"postgres"`
	Log    LogStruct    `yaml:"log"`
}

// ParseYamlFile the config file
func ParseYamlFile(filename string, c *Config) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}
