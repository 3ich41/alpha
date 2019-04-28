package usecases

import "m15.io/alpha/pkg/domain"

type TaskInteractor struct {
	TaskRepository domain.TaskRepository
}

func (interactor *TaskInteractor) CreateAndPublish(mac, sysid string) error {
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
