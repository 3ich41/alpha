package usecases

import "m15.io/alpha/pkg/domain"

type Logger interface {
	Log(message string) error
}

type TaskInteractor struct {
	TaskRepository domain.TaskRepository
	Logger         Logger
}

func (interactor *TaskInteractor) Create(mac, sysid string) error {
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
