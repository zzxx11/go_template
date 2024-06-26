// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go.uber.org/zap"
	"go_template/internal/adapter/controller"
	"go_template/pkg/agin"
)

// Injectors from wire.go:

func NewGinServer(logger *zap.Logger) *agin.Server {
	demoController := controller.NewDemoController(logger)
	v := ProvideController(demoController)
	server := agin.NewServer(logger, v)
	return server
}
