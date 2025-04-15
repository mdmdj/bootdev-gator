package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	DBURL           string `json:"db_url,omitempty"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName string = ".gatorconfig.json"

func getConfigPath() (configPath string, err error) {
	var homeDir string
	homeDir, err = os.UserHomeDir()

	if err != nil {
		//TODO: add log / decorate error
		return
	}

	configPath = path.Join(homeDir, configFileName)
	return
}

func Read() (config *Config, err error) {

	var configFilePath string
	configFilePath, err = getConfigPath()

	if err != nil {
		return
	}

	var configFileBytes []byte
	configFileBytes, err = os.ReadFile(configFilePath)

	if err != nil {
		//TODO: add log / decorate error
		return
	}

	err = json.Unmarshal(configFileBytes, &config)
	if err != nil {
		//TODO: add log / decorate error
		return
	}

	return
}

func (c *Config) SetUser(username string) (err error) {
	c.CurrentUserName = username

	err = c.write()

	return
}

const normalFileMode os.FileMode = 0644

//const normalDirectory os.FileMode = 0755

func (c *Config) write() (err error) {
	var configBytes []byte
	configBytes, err = json.MarshalIndent(*c, "", "\t")

	if err != nil {
		return
	}

	var configFilePath string
	configFilePath, err = getConfigPath()

	if err != nil {
		return
	}

	err = os.WriteFile(configFilePath, configBytes, normalFileMode)
	if err != nil {
		return
	}

	return
}
