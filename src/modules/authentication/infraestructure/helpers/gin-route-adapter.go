package helpers

import (
	"com.task-go-api.com/dudu.com/src/modules/authentication/presentation/dtos"
	"github.com/gin-gonic/gin"
)

func GinRouteAdapter[BODY_T, RESULT any](c *gin.Context, handler func(body BODY_T) (RESULT, *dtos.ErrorDto)) {
	var body BODY_T
	var headers interface{}
	c.BindJSON(&body)
	c.BindHeader(&headers)

	result, err := handler(body)

	if err != nil {
		c.JSON(err.StatusCode, err)
	} else {
		c.JSON(200, result)
	}
}
