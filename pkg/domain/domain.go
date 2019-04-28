package domain

// Task - structure representing switch task sent from client
type Task struct {
	Mac       string
	Sysid     string
	IPAddr    string
	Username  string
	Timestamp string
}

type TaskRepository interface {
	Publish(task Task) error
}
