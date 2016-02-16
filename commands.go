package goci

import "os/exec"

// The DefaultRuncBinary, i.e. 'runc'.
var DefaultRuncBinary = RuncBinary("runc")

// RuncBinary is the path to a runc binary.
type RuncBinary string

// StartCommand creates a start command using the default runc binary name.
func StartCommand(path, id string) *exec.Cmd {
	return DefaultRuncBinary.StartCommand(path, id)
}

// ExecCommand creates an exec command using the default runc binary name.
func ExecCommand(id, processJSONPath string) *exec.Cmd {
	return DefaultRuncBinary.ExecCommand(id, processJSONPath)
}

// KillCommand creates a kill command using the default runc binary name.
func KillCommand(id, signal string) *exec.Cmd {
	return DefaultRuncBinary.KillCommand(id, signal)
}

// StartCommand returns an *exec.Cmd that, when run, will execute a given bundle.
func (runc RuncBinary) StartCommand(path, id string) *exec.Cmd {
	cmd := exec.Command(string(runc), "start", id)
	cmd.Dir = path
	return cmd
}

// ExecCommand returns an *exec.Cmd that, when run, will execute a process spec
// in a running container.
func (runc RuncBinary) ExecCommand(id, processJSONPath string) *exec.Cmd {
	return exec.Command(
		string(runc), "exec", id, "-p", processJSONPath,
	)
}

// KillCommand returns an *exec.Cmd that, when run, will signal the running
// container.
func (runc RuncBinary) KillCommand(id, signal string) *exec.Cmd {
	return exec.Command(
		string(runc), "kill", id, signal,
	)
}
