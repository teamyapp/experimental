package git

import (
	"os"
	"os/exec"
)

type CommandExecutor interface {
	Execute(cmd string, args ...string) (string, error)
}

type ShellExecutor struct {
}

var _ CommandExecutor = (*ShellExecutor)(nil)

func (r ShellExecutor) Execute(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	c.Stderr = os.Stdout
	out, err := c.Output()
	return string(out), err
}
