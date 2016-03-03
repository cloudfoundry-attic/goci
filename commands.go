package goci

import "os/exec"

// The DefaultRuncBinary, i.e. 'runc'.
var DefaultRuncBinary = RuncBinary("runc")

// RuncBinary is the path to a runc binary.
type RuncBinary string

// StartCommand creates a start command using the default runc binary name.
func StartCommand(path, id string, detach bool) *exec.Cmd {
	return DefaultRuncBinary.StartCommand(path, id, detach)
}

// ExecCommand creates an exec command using the default runc binary name.
func ExecCommand(id, processJSONPath, pidFilePath string) *exec.Cmd {
	return DefaultRuncBinary.ExecCommand(id, processJSONPath, pidFilePath)
}

// KillCommand creates a kill command using the default runc binary name.
func KillCommand(id, signal string) *exec.Cmd {
	return DefaultRuncBinary.KillCommand(id, signal)
}

// DeleteCommand creates deletes a container using the default runc binary name.
func DeleteCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.DeleteCommand(id)
}

func EventsCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.EventsCommand(id)
}

// StartCommand returns an *exec.Cmd that, when run, will execute a given bundle.
func (runc RuncBinary) StartCommand(path, id string, detach bool) *exec.Cmd {
	args := []string{"start", id}
	if detach {
		args = []string{"start", "-d", id}
	}

	cmd := exec.Command(string(runc), args...)
	cmd.Dir = path
	return cmd
}

// ExecCommand returns an *exec.Cmd that, when run, will execute a process spec
// in a running container.
func (runc RuncBinary) ExecCommand(id, processJSONPath, pidFilePath string) *exec.Cmd {
	return exec.Command(
		string(runc), "exec", id, "--pid-file", pidFilePath, "-p", processJSONPath,
	)
}

// EventsCommand returns an *exec.Cmd that, when run, will retrieve events for the container
func (runc RuncBinary) EventsCommand(id string) *exec.Cmd {
	return exec.Command(
		string(runc), "events", id,
	)
}

// KillCommand returns an *exec.Cmd that, when run, will signal the running
// container.
func (runc RuncBinary) KillCommand(id, signal string) *exec.Cmd {
	return exec.Command(
		string(runc), "kill", id, signal,
	)
}

// DeleteCommand returns an *exec.Cmd that, when run, will signal the running
// container.
func (runc RuncBinary) DeleteCommand(id string) *exec.Cmd {
	return exec.Command(string(runc), "delete", id)
}
