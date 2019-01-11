package helper

import (
	"github.com/Unknwon/goconfig"
	"errors"
	"sync"
	"webBlog/model"
)

var ConfigManagerIns *ConfigManager

func init() {
	//读取框架运行配置
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

type AppConfig struct {
	configs map[string]string
	mtx sync.Mutex
}
var appConfig *AppConfig
func GetAppConf() *AppConfig{
	if appConfig == nil {
		appConfig = &AppConfig{}
	}
	return appConfig
}

func (a *AppConfig)Get(key string) string {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	value, ok := a.configs[key]
	if ok {
		return value
	}
	return ""
}
func (a *AppConfig)Set(key, value string)  {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.configs[key] = value
}
func (a *AppConfig)Clear()  {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.configs = make(map[string]string)
}
func (a *AppConfig)Init()  {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.configs = model.GetAllConfig()
}
func (a *AppConfig) GetAll() map[string]string  {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	return a.configs
}




