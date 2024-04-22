package main

import (
	"go_template/internal/adapter/controller"
	"go_template/pkg/agin"
)

func ProvideController(demoController *controller.DemoController) []agin.Controller {
	return []agin.Controller{
		demoController,
	}
}
