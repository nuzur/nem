package config

import (
	"fmt"
	"os"
	"strings"

	"errors"
	"go.uber.org/config"
	"path"
)

type DBs []DB

type DB struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pswd"`
	Params   string `yaml:"params"`
	Driver   string `yaml:"driver"`
}

type AWS struct {
	Region string `yaml:"region"`
	KeyID  string `yaml:"key_id"`
	Secret string `yaml:"secret"`
	Bucket string `yaml:"bucket"`
}

func (db DB) Path() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.Params)
}

func New() (config.Provider, error) {
	configPath := os.Getenv("CONFIG")
	envPath := os.Getenv("ENV")
	return NewWithPathAndEnviorment(configPath, envPath)
}

func NewWithPathAndEnviorment(p string, env string) (config.Provider, error) {
	paths := strings.Split(p, ",")
	configFiles := []config.YAMLOption{}
	for _, pt := range paths {
		fi, err := os.Stat(pt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			if env == "" {
				env = "base"
			}
			configFiles = append(configFiles, config.File(path.Join(pt, "base.yaml")))
			envPath := path.Join(pt, fmt.Sprintf("%s.yaml", env))
			if _, err := os.Stat(envPath); err == nil || !errors.Is(err, os.ErrNotExist) {
				configFiles = append(configFiles, config.File(envPath))
			}

		case mode.IsRegular():
			configFiles = append(configFiles, config.File(pt))
		}
	}

	yaml, err := config.NewYAML(
		configFiles...,
	)

	if err != nil {
		fmt.Printf("error loading config: %v", err)
		return nil, err
	}

	return yaml, nil
}
