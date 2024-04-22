package agin

import (
	"github.com/gin-gonic/gin"
	"go_template/configs"
)

type Options struct {
	ListenAddr  string `json:"listen_addr" yaml:"listen_addr" mapstructure:"listen_addr"`
	middlewares []gin.HandlerFunc
}

func NewOptionsFromConfig(config configs.Config) Options {
	options := Options{
		ListenAddr: config.Gin.ListenAddr,
	}

	return options
}

func (opt Options) OptionListenAddr(addr string) Options {
	opt.ListenAddr = addr
	return opt
}

func (opt Options) OptionsMiddlewares(middlewares ...gin.HandlerFunc) Options {
	opt.middlewares = append(opt.middlewares, middlewares...)
	return opt
}
