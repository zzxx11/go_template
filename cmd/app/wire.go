//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"go_template/internal"
	"go_template/pkg/agin"
)

func NewGinServer(logger *zap.Logger) *agin.Server {
	wire.Build(ProvideController,
		internal.Provider,
		agin.Provider,
	)
	return &agin.Server{}
}
