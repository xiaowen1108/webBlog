package helper

import (
	"github.com/Unknwon/goconfig"
	"errors"
)

var ConfigManagerIns *ConfigManager

func init() {
	ConfigManagerIns = &ConfigManager{}
}
type ConfigManager struct {
	Status bool
	ConfigPath string
	Config *goconfig.ConfigFile
}

func InitConfigManager(configPath string) {
	config, err := goconfig.LoadConfigFile(configPath)
	CheckErr(err)
	ConfigManagerIns = &ConfigManager{true,configPath, config}
}
func GetConfig() *ConfigManager{
	return ConfigManagerIns
}
func (c *ConfigManager) GetValue(sec, key string) string {
	if c.Status {
		value, err := c.Config.GetValue(sec, key)
		CheckErr(err)
		return value
	} else {
		panic(errors.New("Config Error"))
	}
}

func (c *ConfigManager) GetSection(sec string) map[string]string {
	if c.Status {
		value, err := c.Config.GetSection(sec)
		CheckErr(err)
		return value
	} else {
		panic(errors.New("Config Error"))
	}
}


