package domain

type Conf struct {
	Username string   `json:"username" binding:"required"`
	Buttons  []Button `json:"buttons" binding:"required"`
}

type Button struct {
	Text  string `json:"text" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ConfRequest struct {
	Username  string `json:"username" binding:"required"`
	IPAddr    string `json:"ipaddr" binding:"required"`
	Mac       string `json:"mac" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

type ConfRepository interface {
	GetConf(confRequest *ConfRequest) (*Conf, error)
}
