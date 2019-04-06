package domain

type Task struct {
	Mac   string
	Sysid string
}

type TaskRepository interface {
	Publish(task Task) error
}
