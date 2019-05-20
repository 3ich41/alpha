package domain

// Conf represents message sent from server to configure client app
type Conf struct {
	Username string   `json:"username" binding:"required"`
	Buttons  []Button `json:"buttons" binding:"required"`
}

// Button is a part of configuration and contains text to display and value to send
type Button struct {
	Text  string `json:"text" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// ConfRequest represents request message sent from client to server
type ConfRequest struct {
	Username  string `json:"username" binding:"required"`
	IPAddr    string `json:"ipaddr" binding:"required"`
	Mac       string `json:"mac" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

// ConfRepository is an interface representing repository containing client app configurations
type ConfRepository interface {
	GetConf(confRequest *ConfRequest) (*Conf, error)
}
