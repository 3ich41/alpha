package interfaces

import "github.com/gin-gonic/gin"

type Task struct {
	Mac   string `json:"mac" binding:"required"`
	SysID string `json:"sysid" binding:"required"`
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
		c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
		return
	}

	if (task.Mac != "") && (task.SysID != "") {
		err := handler.TaskInteractor.CreateAndPublish(task.Mac, task.SysID)
		if err != nil {
			c.JSON(500, gin.H{"Status": "Nie utworzono zadania"})
			return
		}

		c.JSON(200, gin.H{"Status": "OK"})

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
