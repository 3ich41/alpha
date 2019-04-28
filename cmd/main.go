package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"m15.io/alpha/pkg/config"
	"m15.io/alpha/pkg/infrastructure"
	"m15.io/alpha/pkg/interfaces"
	"m15.io/alpha/pkg/usecases"

	log "github.com/sirupsen/logrus"
)

var appName = "alpha"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Infof("Starting service %v...", appName)
	config.InitConfig()
	mqHandler := infrastructure.NewMessagingClient(
		config.Config.MqHostname,
		config.Config.MqPort,
		config.Config.MqUsername,
		config.Config.MqPassword)

	taskInteractor := new(usecases.TaskInteractor)
	taskInteractor.TaskRepository = interfaces.NewMqTaskRepo(mqHandler)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.TaskInteractor = taskInteractor

	engine := gin.Default()
	v1 := engine.Group("api/v1")
	{
		v1.POST("/switch", webserviceHandler.CreateTask)
	}
	engine.Run(":8090")
}
