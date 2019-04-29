package interfaces

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Task struct {
	Mac       string `json:"mac" binding:"required"`
	SysID     string `json:"sysid" binding:"required"`
	IPAddr    string `json:"ipaddr" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

type TaskInteractor interface {
	CreateAndPublish(mac, sysid string) error
}

type WebserviceHandler struct {
	TaskInteractor TaskInteractor
}

func (handler WebserviceHandler) CreateTask(c *gin.Context) {
	var task Task
	err := c.BindJSON(&task)
	if err != nil {
		rawData, _ := ioutil.ReadAll(c.Request.Body)
		log.WithFields(log.Fields{
			"data": string(rawData),
		}).Error("Could not bind json data to Task object. Check if data contains all required fields")
		c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
		return
	}

	if (task.Mac != "") && (task.SysID != "") && (task.IPAddr != "") && (task.Username != "") && (task.Timestamp != "") {
		err := handler.TaskInteractor.CreateAndPublish(task.Mac, task.SysID)
		if err != nil {
			c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
			return
		}

		c.JSON(200, gin.H{"Status": "OK"})

	} else {
		log.WithFields(log.Fields{
			"task": task,
		}).Error("Could not process task. Fields are empty")
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
