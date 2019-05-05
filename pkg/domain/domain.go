package domain

// Task - structure representing switch task sent from client
type Task struct {
	Mac       string `json:"mac" binding:"required"`
	SysID     string `json:"sysid" binding:"required"`
	IPAddr    string `json:"ipaddr" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

type TaskRepository interface {
	Publish(task *Task) error
}
