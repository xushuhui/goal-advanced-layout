package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"

	"fm-suggest/internal/conf"
)

func NewConfig(path string) *conf.Bootstrap {
	var conf conf.Bootstrap

	yamlFile, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("fail to load the config,%s\n", err)
		return nil
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Printf("Serialization configuration failed.%s\n", err)
		return nil
	}
	return &conf
}
