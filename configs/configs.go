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
	Log         Log    `yaml:"log"`
}
type Gin struct {
	ListenAddr string `json:"listen_addr" yaml:"listen_addr" mapstructure:"listen_addr"`
}
type Log struct {
	// Writer is used to specify where to send the logs.
	// default is os.Stdout, which is the standard output.
	// values can be: os.File, os.Stdout, os.Stderr.
	Writer string `yaml:"writer"`
	// Level is the minimum level to log.
	Level string `yaml:"level"`
	// Path is the path to the log file. which is used when Writer is os.File.  if the given path is a directory (the given path ends with a slash),
	// the logger will create a file named with the datetime to write log.
	Path string `yaml:"path"`
	// Format is the log format. value is json or text.
	Format string `yaml:"format"`
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
