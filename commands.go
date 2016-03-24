package goci

import "os/exec"

// The DefaultRuncBinary, i.e. 'runc'.
var DefaultRuncBinary = RuncBinary{RuncBinary: "runc"}

// RuncBinary is the path to a runc binary.
type RuncBinary struct {
	logFile    string
	RuncBinary string
}

func NewBinary(name string) RuncBinary {
	return RuncBinary{
		logFile:    "/dev/null",
		RuncBinary: name,
	}
}

type Runc interface {
	StartCommand(path, id string, detach bool) *exec.Cmd
	ExecCommand(id, processJSONPath, pidFilePath string) *exec.Cmd
	KillCommand(id, signal string) *exec.Cmd
	StateCommand(id string) *exec.Cmd
	StatsCommand(id string) *exec.Cmd
	DeleteCommand(id string) *exec.Cmd
	EventsCommand(id string) *exec.Cmd
}

type RuncCmd interface {
	WithLogFile(file string) *exec.Cmd
}

// WithLogFile returns a runc binary whose methods all log to the given log file
func (runc RuncBinary) WithLogFile(file string) Runc {
	return RuncBinary{
		logFile:    file,
		RuncBinary: runc.RuncBinary,
	}
}

func WithLogFile(file string) Runc {
	return DefaultRuncBinary.WithLogFile(file)
}

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

// StateCommands creates a command that gets the state of a container using the default runc binary name.
func StateCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.StateCommand(id)
}

// StatsCommands creates a command that gets the stats of a container using the default runc binary name.
func StatsCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.StatsCommand(id)
}

// DeleteCommand creates a command that deletes a container using the default runc binary name.
func DeleteCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.DeleteCommand(id)
}

func EventsCommand(id string) *exec.Cmd {
	return DefaultRuncBinary.EventsCommand(id)
}

// StartCommand returns an *exec.Cmd that, when run, will execute a given bundle.
func (runc RuncBinary) StartCommand(path, id string, detach bool) *exec.Cmd {
	args := []string{"--debug", "--log", runc.logFile, "start"}
	if detach {
		args = append(args, "-d")
	}

	args = append(args, id)

	cmd := exec.Command(runc.RuncBinary, args...)
	cmd.Dir = path
	return cmd
}

// ExecCommand returns an *exec.Cmd that, when run, will execute a process spec
// in a running container.
func (runc RuncBinary) ExecCommand(id, processJSONPath, pidFilePath string) *exec.Cmd {
	return exec.Command(
		runc.RuncBinary, "--debug", "--log", runc.logFile, "exec", id, "--pid-file", pidFilePath, "-p", processJSONPath,
	)
}

// EventsCommand returns an *exec.Cmd that, when run, will retrieve events for the container
func (runc RuncBinary) EventsCommand(id string) *exec.Cmd {
	return exec.Command(
		runc.RuncBinary, "--debug", "--log", runc.logFile, "events", id,
	)
}

// KillCommand returns an *exec.Cmd that, when run, will signal the running
// container.
func (runc RuncBinary) KillCommand(id, signal string) *exec.Cmd {
	return exec.Command(
		runc.RuncBinary, "--debug", "--log", runc.logFile, "kill", id, signal,
	)
}

// StateCommand returns an *exec.Cmd that, when run, will get the state of the
// container.
func (runc RuncBinary) StateCommand(id string) *exec.Cmd {
	return exec.Command(runc.RuncBinary, "--debug", "--log", runc.logFile, "state", id)
}

// StatsCommand returns an *exec.Cmd that, when run, will get the stats of the
// container.
func (runc RuncBinary) StatsCommand(id string) *exec.Cmd {
	return exec.Command(runc.RuncBinary, "--debug", "--log", runc.logFile, "events", "--stats", id)
}

// DeleteCommand returns an *exec.Cmd that, when run, will signal the running
// container.
func (runc RuncBinary) DeleteCommand(id string) *exec.Cmd {
	return exec.Command(runc.RuncBinary, "--debug", "--log", runc.logFile, "delete", id)
}
