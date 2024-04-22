package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectName string `yaml:"project_name"`
	Env         string `yaml:"env"`
	Gin         Gin    `yaml:"gin"`
}
type Gin struct {
	ListenAddr string `json:"listen_addr" yaml:"listen_addr" mapstructure:"listen_addr"`
}

var config = new(Config)

func LoadConfig(path string) *Config {

	if len(path) == 0 {
		panic("config file path is empty")
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(config); err != nil {
		panic(err)
	}
	return config
}

func Get() *Config {
	return config
}
