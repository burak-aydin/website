package pack

import (
	"github.com/burak-aydin/website/email"
	"github.com/burak-aydin/website/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RegisterRouters(e *gin.Engine) {
	e.GET("/api/ping", func(context *gin.Context) {
		log.Info("ping/pong")
		context.JSON(http.StatusOK, "pong")
	})
	e.POST("/api/send-email", func(context *gin.Context) {
		var params SendEmailParams
		if err := utils.HandleError(context.ShouldBindJSON(&params)); err != nil {
			log.WithError(err).Info("binding error for send-email")
			context.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorJson{Error: err.Error()})
			return
		}
		email.SendEmail(params.Content)
		context.JSON(http.StatusOK, "Success")
	})
}
