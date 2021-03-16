package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Username string
	ApiKey   string
	SavePath string
	Posthook string
	Height   uint
	Width    uint
	Interval uint
}

var lfmConfig Config

func parseConfig(data string) (*Config, error) {
	config := new(Config)
	conf, err := toml.Load(data)

	if err != nil {
		return nil, err
	}

	username := conf.Get("USERNAME")
	if username == nil {
		return nil, errors.New("[CONFIG] USERNAME not defined")
	}

	apiKey := conf.Get("API_KEY")
	if apiKey == nil {
		return nil, errors.New("[CONFIG] API_KEY not defined")
	}

	savePath := conf.Get("SAVE_PATH")
	if savePath == nil {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		config.SavePath = fmt.Sprintf("%s/.lfmbg.png", home)
	} else {
		config.SavePath = savePath.(string)
	}

	width := conf.Get("WIDTH")
	if width == nil {
		return nil, errors.New("[CONFIG] WIDTH not defined")
	}

	height := conf.Get("HEIGHT")
	if height == nil {
		return nil, errors.New("[CONFIG] HEIGHT not defined")
	}

	interval := conf.Get("INTERVAL")
	if interval == nil {
		config.Interval = 5
	} else {
		config.Interval = uint(interval.(int64))
	}

	posthook := conf.Get("POSTHOOK")
	if posthook == nil {
		config.Posthook = ""
	} else {
		config.Posthook = posthook.(string)
	}

	config.Username = username.(string)
	config.ApiKey = apiKey.(string)
	config.Width = uint(width.(int64))
	config.Height = uint(height.(int64))

	return config, nil
}

func readConfig() (*Config, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(home + "/.config/lfmbg/config")

	if err != nil {
		return nil, err
	}

	return parseConfig(string(dat))

}

func GetConfig() (*Config, error) {
	config, err := readConfig()

	if err != nil {
		return nil, err
	}

	return config, nil

}
