package main

import (
	"github.com/burak-aydin/website/pack"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func GetHandler() *gin.Engine {
	gin.SetMode("debug") //debug or release
	engine := gin.New()
	pack.RegisterRouters(engine)
	return engine
}

func main() {
	log.Info("start-up")

	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      GetHandler(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}

	log.Info("starting server")
	err := s.ListenAndServe()
	if err != nil {
		log.WithError(err).Error("ListenAndServe returned err")
	}

	log.Info("exit")
}
