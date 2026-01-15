package devboard

import "time"

type Project struct {
	Id        int
	Path      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Project) Title() string { return p.Name }
func (p Project) Description() string { return p.Path }
func (p Project) FilterValue() string { return p.Name }
