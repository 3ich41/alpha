package domain

// Task - structure representing switch task sent from client
type Task struct {
	Mac       string `json:"mac" binding:"required"`
	SysID     string `json:"sysid" binding:"required"`
	IPAddr    string `json:"ipaddr" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

// TaskRepository is an interface representing repository where requests to configure clients are written to
type TaskRepository interface {
	Publish(task *Task) error
}
