package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_template/pkg/agin"
	"go_template/pkg/framework"
)

type DemoController struct {
	logger *zap.Logger
}

func NewDemoController(logger *zap.Logger) *DemoController {

	return &DemoController{logger: logger}
}

func (ctrl *DemoController) InitRouter(r *gin.Engine) {
	r.GET("/api/hello/v1", ctrl.get)
	r.GET("/api/hello/v1/err", ctrl.err)
}

func (ctrl *DemoController) get(c *gin.Context) {
	agin.WriteResponse(c, agin.ResponseData{Data: map[string]string{"msg": "hello world!"}}, nil)
}

func (ctrl *DemoController) err(c *gin.Context) {
	err := agin.NewHTTPError(framework.CommonInternalErr.New("common error").(framework.AppError), 500)
	agin.WriteResponse(c, agin.ResponseData{Data: map[string]string{"msg": "hello world!"}}, err)
}
