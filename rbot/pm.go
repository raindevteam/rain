package rbot

import (
	"errors"
	"os/exec"
	"sync"
)

type Result struct {
	Output string
	Err    error
}

// A ProcessManager handles an individual process of a module. The ProcessManager can recompile and
// invoke a module as well as terminate its process. Mostly used by the bot for module management.
type ProcessManager struct {
	Name        string
	Type        string
	Path        string
	cmd         *exec.Cmd
	running     bool
	processDone chan *Result
	lastResult  *Result
	mu          sync.Mutex
}

// NewProcessManager will return a ProcessManager of the correct type for a given module name and
// its path.
func NewProcessManager(name string, cmdtype string, path string) *ProcessManager {
	pm := &ProcessManager{
		Name:    name,
		Type:    cmdtype,
		Path:    path,
		running: false,
		mu:      sync.Mutex{},
	}
	return pm
}

func (pm *ProcessManager) runCommand(name string, args ...string) chan *Result {
	pm.mu.Lock()
	pm.running = true
	pm.mu.Unlock()

	ch := make(chan *Result)

	go func(ch chan *Result, pm *ProcessManager) {
		pm.cmd = exec.Command(name, args...)
		output, err := pm.cmd.CombinedOutput()
		s := string(output[:])

		pm.mu.Lock()

		pm.lastResult = &Result{s, err}
		pm.running = false

		pm.mu.Unlock()

		ch <- pm.lastResult
	}(ch, pm)

	return ch
}

func (pm *ProcessManager) WaitForCommand() *Result {
	pm.mu.Lock()
	if !pm.running {
		return nil
	}
	pm.mu.Unlock()
	return <-pm.processDone
}

func (pm *ProcessManager) LastResult() *Result {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	return pm.lastResult
}

// Recompile will attempt to compile any source code of a module if needed. If compilation fails,
// an error is returned. Nothing will be done for types that do not require compilation.
func (pm *ProcessManager) Recompile() error {
	switch pm.Type {
	case "go":
		cmd := pm.runCommand("go", "install", pm.Path+"/"+pm.Name)
		res := <-cmd
		if res.Err != nil {
			return errors.New("Could not recompile")
		}
	}
	return nil
}

// Start creates a new exec.Command and stores it. Will return an error if the Command fails to
// start.
// TODO: Cleanup and Refactor
func (pm *ProcessManager) Start() chan *Result {
	switch pm.Type {
	case "js":
		return pm.runCommand("node", pm.Path+"/"+pm.Name)
	case "go":
		return pm.runCommand(pm.Name)
	case "py":
		return pm.runCommand("python", pm.Path+"/"+pm.Name)
	default:
		// Keep in mind that the will make it virtually impossible to reach here
		return nil
	}
}

// Kill attempts to kill its Cmd process, an error is returned if a failure occurs.
func (pm *ProcessManager) Kill() error {
	err := pm.cmd.Process.Kill()
	pm.running = false
	if err != nil {
		return err
	}
	return nil
}
