package interfaces

import (
	"encoding/json"

	"m15.io/alpha/pkg/domain"
)

type MqHandler interface {
	PublishOnQueue(msg []byte, queueName string) error
}

type MqRepo struct {
	mqHandler MqHandler
}

type MqTaskRepo MqRepo

func NewMqTaskRepo(mqHandler MqHandler) *MqTaskRepo {
	mqTaskRepo := new(MqTaskRepo)
	mqTaskRepo.mqHandler = mqHandler
	return mqTaskRepo
}

func (repo *MqTaskRepo) Publish(task domain.Task) error {
	msg, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = repo.mqHandler.PublishOnQueue(msg, "unpreparedTasks")
	if err != nil {
		return err
	}

	return nil
}
