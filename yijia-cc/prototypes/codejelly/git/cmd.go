package git

import "os/exec"

type CommandExecutor interface {
	Execute(cmd string, args ...string) (string, error)
}

type ShellExecutor struct {
}

var _ CommandExecutor = (*ShellExecutor)(nil)

func (r ShellExecutor) Execute(cmd string, args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	return string(out), err
}
