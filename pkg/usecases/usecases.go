package usecases

import (
	"m15.io/alpha/pkg/domain"

	log "github.com/sirupsen/logrus"
)

type TaskInteractor struct {
	TaskRepository domain.TaskRepository
}

func (interactor *TaskInteractor) CreateAndPublish(mac, sysid string) error {
	log.Debug("Creating and publishing task")
	task := domain.Task{
		Mac:   mac,
		Sysid: sysid,
	}

	err := interactor.TaskRepository.Publish(task)
	if err != nil {
		return err
	}

	return nil
}
