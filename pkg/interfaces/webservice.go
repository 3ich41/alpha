package interfaces

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"m15.io/alpha/pkg/domain"
)

type TaskInteractor interface {
	Publish(task *domain.Task) error
}

type ConfInteractor interface {
	ConfigureClientApp(confRequest *domain.ConfRequest) (*domain.Conf, error)
}

type WebserviceHandler struct {
	TaskInteractor TaskInteractor
	ConfInteractor ConfInteractor
}

func (handler *WebserviceHandler) NewTask(c *gin.Context) {
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
		c.JSON(400, gin.H{"Status": "Fields are empty"})
	}
}

func (handler *WebserviceHandler) ConfigureClient(c *gin.Context) {
	log.Debug("Received request to configure client")

	confRequest, err := handler.getConfRequest(c)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Could not bind json data to Conf object. Check if data contains all required fields")

		c.JSON(500, gin.H{"Status": "Could not bind json data to Conf object. Check if data contains all required fields"})
		return
	}

	handler.handleConfRequest(confRequest, c)
}

func (handler *WebserviceHandler) getConfRequest(c *gin.Context) (*domain.ConfRequest, error) {
	confRequest := new(domain.ConfRequest)

	err := c.BindJSON(confRequest)
	if err != nil {
		return nil, err
	}

	return confRequest, nil
}

func (handler *WebserviceHandler) handleConfRequest(confRequest *domain.ConfRequest, c *gin.Context) {
	if (confRequest.Mac != "") && (confRequest.IPAddr != "") && (confRequest.Username != "") && (confRequest.Timestamp != "") {
		handler.handleCorrectConfRequest(confRequest, c)
	} else {
		handler.handleEmptyFieldsConfRequest(confRequest, c)
	}
}

func (handler *WebserviceHandler) handleCorrectConfRequest(confRequest *domain.ConfRequest, c *gin.Context) {
	conf, err := handler.ConfInteractor.ConfigureClientApp(confRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"err":       err.Error(),
			"Mac":       confRequest.Mac,
			"IPAddr":    confRequest.IPAddr,
			"Username":  confRequest.Username,
			"Timestamp": confRequest.Timestamp,
		}).Error("Failed processing configuration request")

		c.JSON(500, gin.H{"Status": "Failed processing configuration request"})
		return
	}
	handler.sendConfToClientApp(conf, confRequest, c)
}

func (handler *WebserviceHandler) handleEmptyFieldsConfRequest(confRequest *domain.ConfRequest, c *gin.Context) {
	log.WithFields(log.Fields{
		"Mac":       confRequest.Mac,
		"IPAddr":    confRequest.IPAddr,
		"Username":  confRequest.Username,
		"Timestamp": confRequest.Timestamp,
	}).Error("Could not process configuration request. Fields are empty")
	c.JSON(400, gin.H{"Status": "Could not process configuration request. Fields are empty"})
}

func (handler *WebserviceHandler) sendConfToClientApp(conf *domain.Conf, confRequest *domain.ConfRequest, c *gin.Context) {
	log.WithFields(log.Fields{
		"Conf":      conf,
		"Mac":       confRequest.Mac,
		"IPAddr":    confRequest.IPAddr,
		"Username":  confRequest.Username,
		"Timestamp": confRequest.Timestamp,
	}).Info("Sending configuration to client application")
	c.JSON(200, conf)
}
