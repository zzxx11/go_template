package internal

import (
	"github.com/google/wire"
	"go_template/internal/adapter/controller"
)

var Provider = wire.NewSet(
	controller.NewDemoController,

	//repository.NewDemoRepository,
)
