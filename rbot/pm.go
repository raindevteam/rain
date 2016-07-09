package rbot

import (
	"errors"
	"os/exec"
)

// A ProcessManager handles an individual process of a module. The ProcessManager can recompile and
// invoke a module as well as terminate its process. Mostly used by the bot for module management.
type ProcessManager struct {
	Name string
	Type string
	Path string
	Cmd  *exec.Cmd
}

// NewProcessManager will return a ProcessManager of the correct type for a given module name and
// its path.
func NewProcessManager(name string, cmdtype string, path string) *ProcessManager {
	pm := &ProcessManager{
		Name: name,
		Type: cmdtype,
		Path: path,
	}
	return pm
}

// Recompile will attempt to compile any source code of a module if needed. If compilation fails,
// an error is returned. Nothing will be done for types that do not require compilation.
func (c *ProcessManager) Recompile() error {
	switch c.Type {
	case "go":
		output, err := exec.Command("go", "install", c.Path+"/"+c.Name).CombinedOutput()
		s := string(output[:])
		if err != nil {
			return errors.New("Could not recompile: " + s)
		}
	default:
	}

	return nil
}

// Start creates a new exec.Command and stores it. Will return an error if the Command fails to
// start.
// TODO: Cleanup and Refactor
func (c *ProcessManager) Start() (err error) {
	switch c.Type {
	case "js":
		c.Cmd = exec.Command("node", c.Path+"/"+c.Name)
		err = c.Cmd.Start()
	case "go":
		c.Cmd = exec.Command(c.Name)
		err = c.Cmd.Start()
	case "py":
		c.Cmd = exec.Command("python", c.Path+"/"+c.Name)
		err = c.Cmd.Start()
	}

	if err != nil {
		return err
	}

	return nil
}

// Kill attempts to kill its Cmd process, an error is returned if a failure occurs.
func (c *ProcessManager) Kill() error {
	err := c.Cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}
