package config

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-archaius/sources/file-source"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

var configurations *Config

func Init(file string) error {
	if err := archaius.AddFile(file, archaius.WithFileHandler(filesource.UseFileNameAsKeyContentAsValue)); err != nil {
		return err
	}
	_, filename := filepath.Split(file)
	content := archaius.GetString(filename, "")
	configurations = &Config{}
	if err := yaml.Unmarshal([]byte(content), configurations); err != nil {
		return err
	}
	return nil
}

func GetDB() DB {
	return configurations.DB
}
