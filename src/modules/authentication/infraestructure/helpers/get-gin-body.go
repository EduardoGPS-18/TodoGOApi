package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetGinBody[REQUEST_BODY any](ctx *gin.Context) REQUEST_BODY {
	var body REQUEST_BODY

	ctx.BindJSON(&body)

	return body
}
