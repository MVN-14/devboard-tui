package devboard

import (
	"encoding/json"
	"time"
)

type Project struct {
	Id       int
	Command  string
	Name     string
	Path     string
	OpenedAt *time.Time `json:"opened_at"`
}

func (p Project) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}
