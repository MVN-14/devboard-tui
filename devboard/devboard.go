package devboard

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

func Devboard(args ...string) ([]byte, error) {
	cmd := exec.Command("devboard", args...)

	out, err := cmd.Output()
	if err != nil {
		return []byte{}, fmt.Errorf("devboard %s error:%s %s", args, out, cmd.Stderr)
	}

	return out, nil
}

func OpenProject(id int) error {
	_, err := Devboard("open", strconv.Itoa(id))
	if err != nil {
		return err
	}

	return nil
}

func LoadProjects() ([]Project, error) {
	out, err := Devboard("list")
	if err != nil {
		return []Project{}, err
	}
	projects := []Project{}

	if err := json.Unmarshal(out, &projects); err != nil {
		return []Project{}, fmt.Errorf("Error parsing output from list: %s", err)
	}
	
	return projects, nil
}

func UpdateProject(p Project) error {
	_, err := Devboard("update", p.String())
	if err != nil {
		return err
	}
	return nil
}

func AddProject(p Project) error {
	_, err := Devboard("add", p.String())
	if err != nil {
		return err
	}
	return nil
}

func DeleteProject(id int) error {
	_, err := Devboard("delete", strconv.Itoa(id))
	if err != nil {
		return err
	}
	return nil
}
