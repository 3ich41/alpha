package interfaces

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"m15.io/alpha/pkg/domain"
)


type TaskInteractor interface {
	Publish(task *domain.Task) error
}

type WebserviceHandler struct {
	TaskInteractor TaskInteractor
}

func (handler WebserviceHandler) NewTask(c *gin.Context) {
	log.Debug("Received request to switch client")
	task := new(domain.Task)

	err := c.BindJSON(task)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Could not bind json data to Task object. Check if data contains all required fields")

		c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
		return
	}

	if (task.Mac != "") && (task.SysID != "") && (task.IPAddr != "") && (task.Username != "") && (task.Timestamp != "") {
		err := handler.TaskInteractor.Publish(task)
		if err != nil {
			log.WithFields(log.Fields{
				"err":       err.Error(),
				"Mac":       task.Mac,
				"SysID":     task.SysID,
				"IPAddr":    task.IPAddr,
				"Username":  task.Username,
				"Timestamp": task.Timestamp,
			}).Error("Could not publish task")

			c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
			return
		}

		log.WithFields(log.Fields{
			"Mac":       task.Mac,
			"SysID":     task.SysID,
			"IPAddr":    task.IPAddr,
			"Username":  task.Username,
			"Timestamp": task.Timestamp,
		}).Info("Task published")
		c.JSON(200, gin.H{"Status": "OK"})

	} else {
		log.WithFields(log.Fields{
			"Mac":       task.Mac,
			"SysID":     task.SysID,
			"IPAddr":    task.IPAddr,
			"Username":  task.Username,
			"Timestamp": task.Timestamp,
		}).Error("Could not process task. Fields are empty")
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
