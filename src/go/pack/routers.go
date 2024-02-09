package pack

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouters(e *gin.Engine) {
	e.GET("/api/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
}
