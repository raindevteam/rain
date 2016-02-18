package rainbot

import (
	"errors"
	"fmt"
	"os/exec"
)

type Commander struct {
	Name string
	Type string
	Path string
	Cmd  *exec.Cmd
}

func NewCommander(name string, cmdtype string, path string) *Commander {
	commander := &Commander{
		Name: name,
		Type: cmdtype,
		Path: path,
	}
	return commander
}

func (c *Commander) Recompile() error {
	switch c.Type {
	case "js":
		// No need to recompile js
	case "go":
		output, err := exec.Command("go", "install", c.Path+"/"+c.Name).CombinedOutput()
		s := string(output[:])
		if err != nil {
			fmt.Println(s)
			return errors.New("Could not recompile")
		}
	}

	return nil
}

func (c *Commander) Start() (err error) {
	switch c.Type {
	case "js":
		c.Cmd = exec.Command("node " + c.Path + "/" + c.Name)
		err = c.Cmd.Start()
	case "go":
		c.Cmd = exec.Command(c.Name)
		err = c.Cmd.Start()
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *Commander) Kill() error {
	err := c.Cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}
