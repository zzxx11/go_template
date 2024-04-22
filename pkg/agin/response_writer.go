package agin

import (
	"github.com/gin-gonic/gin"
	"go_template/pkg/framework"
)

type ResponseData struct {
	Code    framework.ErrorType `json:"code"`
	Message string              `json:"message"`
	Data    any                 `json:"data"`
}
type ListResponseData struct {
	Total int `json:"total"`
	Data  any `json:"data"`
}
type RewardResponseData struct {
	ReceivedReward bool `json:"received_reward"`
	Data           any  `json:"data"`
}

func WriteResponse(c *gin.Context, response ResponseData, err error) {
	if err != nil {
		if httpErr, ok := err.(*HTTPError); ok {
			c.JSON(httpErr.code, ResponseData{
				Code:    framework.GetType(httpErr.AppError),
				Message: httpErr.AppError.Error(),
			})
			//observability.AppendEvents(c.Request.Context(),
			//	"biz.code", framework.GetType(err),
			//	"err.msg", httpErr.AppError.Error(),
			//)
		} else if appErr, ok := err.(framework.AppError); ok {
			c.JSON(500, ResponseData{
				Code:    framework.GetType(err),
				Message: appErr.Error(),
			})
			//observability.AppendEvents(c.Request.Context(),
			//	"biz.code", framework.GetType(err),
			//	"err.msg", appErr.Error(),
			//)
		} else {
			c.JSON(500, ResponseData{
				Code:    framework.CommonInternalErr,
				Message: err.Error(),
			})
			//observability.AppendEvents(c.Request.Context(),
			//	"biz.code", framework.CommonInternalErr,
			//	"err.msg", err.Error(),
			//)
		}
	} else {
		c.JSON(200, response)
	}
}
