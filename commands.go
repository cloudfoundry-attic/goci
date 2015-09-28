package goci

import "os/exec"

// The DefaultRuncBinary, i.e. 'runc'
var DefaultRuncBinary = RuncBinary("runc")

// RuncBinary is the path to a runc binary
type RuncBinary string

// StartCommand creates a start command using the default runc binary name
func StartCommand(path, id string) *exec.Cmd {
	return DefaultRuncBinary.StartCommand(path, id)
}

// ExecCommand creates an exec command using the default runc binary name
func ExecCommand(id, processJSONPath string) *exec.Cmd {
	return DefaultRuncBinary.ExecCommand(id, processJSONPath)
}

// StartCommand returns an *exec.Cmd that, when run, will execute a given bundle
func (runc RuncBinary) StartCommand(path, id string) *exec.Cmd {
	return &exec.Cmd{
		Path: string(runc),
		Args: []string{string(runc), "--id", id, "start", path},
	}
}

// ExecCommand returns an *exec.Cmd that, when run, will execute a process spec
// in a running container
func (runc RuncBinary) ExecCommand(id, processJSONPath string) *exec.Cmd {
	return &exec.Cmd{
		Path: string(runc),
		Args: []string{string(runc), "--id", id, "exec", processJSONPath},
	}
}
