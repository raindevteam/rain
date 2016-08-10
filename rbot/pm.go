// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

package rbot

import (
	"bytes"
	"os/exec"
	"runtime"
	"sync"
)

// The Result struct is used to hold information about ran commands, this makes it easier to debug
// and manage processes.
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

// runCommand  will start a given command (usually a module process). It will also run a goroutine
// that will listen on the kill chan, which when fulfilled, will terminate the current running
// process. Stderr and Stdout are handled manually rather than calling CombinedOutput so that the
// Process Manager can hold a lock up until the command has started.
func (pm *ProcessManager) runCommand(name string, args ...string) *Result {
	pm.mu.Lock()

	pm.processDone = make(chan *Result, 1)
	pm.running = true
	pm.cmd = exec.Command(name, args...)

	// Handle a kill signal
	go func(pm *ProcessManager) {
		<-pm.kill
		pm.mu.Lock()
		pm.cmd.Process.Kill()
		pm.mu.Unlock()
	}(pm)

	var (
		b   bytes.Buffer
		err error
		res *Result
	)

	pm.cmd.Stdout = &b
	pm.cmd.Stderr = &b

	err = pm.cmd.Start()
	if err != nil {
		pm.mu.Unlock()
		return &Result{"", err}
	}

	pm.mu.Unlock()

	err = pm.cmd.Wait()
	output := string(b.Bytes()[:])

	res = &Result{output, err}

	pm.mu.Lock()

	pm.lastResult = res
	pm.running = false

	pm.mu.Unlock()

	pm.processDone <- res
	return res
}

// Wait will wait on the processDone channel, which if fulfilled when a command has finished
// executing. A Result struct will be returned.
func (pm *ProcessManager) Wait() chan *Result {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	if !pm.running {
		return nil
	}

	return pm.processDone
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

// Recompile will attempt to compile any source code of a module if needed. A result struct will be
// returned to the caller for inspection To check if there was an error in compilation.
func (pm *ProcessManager) Recompile() *Result {
	var (
		res *Result
	)

	switch pm.Type {
	case "go":
		res = pm.runCommand("go", "install", pm.Path+"/"+pm.Name)
	default:
		return nil
	}

	return res
}

// Start will run a module process and pass a given port to it via command line arguments. A result
// struct will be returned when module process has exited.
func (pm *ProcessManager) Start(port string) *Result {
	switch pm.Type {
	case "js":
		return pm.runCommand("node", pm.Path+"/"+pm.Name, port)
	case "go":
		return pm.runCommand(pm.Name, port)
	case "py":
		var python = "python"
		if runtime.GOOS != "windows" {
			python = "python3"
		}
		return pm.runCommand(python, pm.Path+"/"+pm.Name, port)
	default:
		// Keep in mind that the bot will make it virtually impossible to reach here
		// But you know how the story goes...
		return nil
	}
}

// Kill will fulfill the kill chan, and any running module process will be terminated.
func (pm *ProcessManager) Kill() {
	pm.kill <- true
}
