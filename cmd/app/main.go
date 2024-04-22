package main

import (
	"go.uber.org/zap"
	"go_template/configs"
	"go_template/global"
	"go_template/pkg/agin"
	"go_template/pkg/framework"
	"go_template/pkg/logger"
)

func main() {
	config := configs.LoadConfig("./configs/configs.yaml")

	var err error
	global.Logger, err = logger.NewJSONLogger(
		//logger.WithDisableConsole(),
		logger.WithField(config.ProjectName, config.Env),
		logger.WithFileRotationP("./logs/log.txt"),
	)
	if err != nil {
		panic(err)
	}

	global.Logger.Info("Starting server...")

	server := NewGinServer(global.Logger).WithOptions(agin.NewOptionsFromConfig(*config).OptionListenAddr(":80"))

	app := framework.NewApp(global.Logger)
	app.WithServer(server)
	err = app.Start()
	if err != nil {
		global.Logger.Error("app exited", zap.Error(err))
	}

}
