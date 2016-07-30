package rbot

import (
	"os/exec"
	"runtime"
	"sync"
)

// The Result struct is used to hold information about ran commands, this makes it easier to debug and manage processes.
type Result struct {
	Output string
	Err    error
}

// A ProcessManager handles an individual process of a module. The ProcessManager can recompile and invoke a module as
// well as terminate its process. Mostly used by the bot for module management.
type ProcessManager struct {
	Name        string
	Type        string
	Path        string
	cmd         *exec.Cmd
	running     bool
	processDone chan *Result
	kill        chan bool
	lastResult  *Result
	mu          sync.RWMutex
}

// NewProcessManager will return a ProcessManager of the correct type for a given module.
func NewProcessManager(name string, cmdtype string, path string) *ProcessManager {
	pm := &ProcessManager{
		Name:    name,
		Type:    cmdtype,
		Path:    path,
		running: false,
		kill:    make(chan bool, 1),
		mu:      sync.RWMutex{},
	}
	return pm
}

// runCommand is the workhorse of a process manager. It will run a command with given arguments in a goroutine. Another
// goroutine is used to wait upon two outcomes:
//
// (A) When the command has finished
// (B) When the command is killed
//
// When the command has finished, lastResult is updated. The caller receives a channel that returns this result when
// the command has exited via one of the outcomes mentioned.
func (pm *ProcessManager) runCommand(name string, args ...string) chan *Result {
	pm.mu.Lock()
	pm.running = true
	pm.cmd = exec.Command(name, args...)
	pm.mu.Unlock()

	ret := make(chan *Result, 1)
	done := make(chan *Result, 1)

	go func(done chan *Result, pm *ProcessManager) {
		pm.mu.Lock()
		output, err := pm.cmd.CombinedOutput()
		pm.mu.Unlock()
		s := string(output[:])

		res := &Result{s, err}
		done <- res

		pm.mu.Lock()

		pm.lastResult = res
		pm.running = false

		pm.mu.Unlock()
	}(done, pm)

	go func(ret chan *Result, done chan *Result, pm *ProcessManager) {
		select {
		case res := <-done:
			ret <- res
		case <-pm.kill:
			pm.mu.RLock()
			pm.cmd.Process.Release()
			pm.cmd.Process.Kill()
			pm.mu.RUnlock()
			res := <-done
			ret <- res
		}
	}(ret, done, pm)

	return ret
}

// WaitForCommand will wait on the processDone channel, which if fullfilled when a command has finished executing. A
// Result struct will be returned.
func (pm *ProcessManager) WaitForCommand() *Result {
	pm.mu.RLock()
	if !pm.running {
		return nil
	}
	pm.mu.RUnlock()
	return <-pm.processDone
}

// IsRunning returns a bool indicating whether a command is running or not.
func (pm *ProcessManager) IsRunning() bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.running
}

// LastResult retrieves the last result struct from the most recent finished command.
func (pm *ProcessManager) LastResult() *Result {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.lastResult
}

// Recompile will attempt to compile any source code of a module if needed. A result struct will be returned to the
// caller for inspection To check if there was an error in compilation).
func (pm *ProcessManager) Recompile() *Result {
	var (
		res *Result
		cmd chan *Result
	)

	switch pm.Type {
	case "go":
		cmd = pm.runCommand("go", "install", pm.Path+"/"+pm.Name)
	default:
		return nil
	}

	res = <-cmd
	return res
}

// Start will run a command via runCommand and return it's channel for the caller.
func (pm *ProcessManager) Start() chan *Result {
	switch pm.Type {
	case "js":
		return pm.runCommand("node", pm.Path+"/"+pm.Name)
	case "go":
		return pm.runCommand(pm.Name)
	case "py":
		var python = "python"
		if runtime.GOOS != "windows" {
			python = "python3"
		}
		return pm.runCommand(python, pm.Path+"/"+pm.Name)
	default:
		// Keep in mind that the bot will make it virtually impossible to reach here
		// But you know how the story goes...
		return nil
	}
}

// Kill will fulfill the kill chan, and any running command will be terminated.
func (pm *ProcessManager) Kill() {
	pm.kill <- true
}
