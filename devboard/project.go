package devboard

import "encoding/json"

type Project struct {
	Id      int
	Command string
	Name    string
	Path    string
}

func (p Project) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}
