package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"m15.io/alpha/pkg/config"
	"m15.io/alpha/pkg/infrastructure"
	"m15.io/alpha/pkg/interfaces"
	"m15.io/alpha/pkg/usecases"
)

func main() {
	config.InitConfig()
	mqHandler, err := infrastructure.NewRabbitMqHandler(
		config.Config.MqUsername,
		config.Config.MqPassword,
		config.Config.MqHostname,
		config.Config.MqPort)
	if err != nil {
		log.Fatal(err)
	}

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
