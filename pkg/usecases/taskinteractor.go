package usecases

import (
	"m15.io/alpha/pkg/domain"
)

// TaskInteractor provides implementation for various use cases
type TaskInteractor struct {
	TaskRepository domain.TaskRepository
}

// Publish publishes task on 'unpreparedTasks' queue in RabbitMQ
func (interactor *TaskInteractor) Publish(task *domain.Task) error {

	err := interactor.TaskRepository.Publish(task)
	if err != nil {
		return err
	}

	return nil
}
