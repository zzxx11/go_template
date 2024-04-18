package main

import (
	"go_template/configs"
	"go_template/pkg/logger"
)

func main() {
	config := configs.LoadConfig("./configs.yaml")
	log, err := logger.NewJSONLogger(
		//logger.WithDisableConsole(),
		logger.WithField(config.ProjectName, config.Env),
		logger.WithFileRotationP("./logs/log.txt"),
	)
	if err != nil {
		panic(err)
	}
	log.Error("message")

}
