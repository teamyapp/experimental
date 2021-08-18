package git

import (
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

type Repository struct {
	rootPath        string
	commandExecutor CommandExecutor
}

func (r Repository) GetFileDiffsBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiff, error) {
	output, err := r.executeGitCommand("diff", "--name-status", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	fileDiffs := make([]entity.FileDiff, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fileDiff, err := newFileDiffFromLine(line)
		if err != nil {
			return nil, err
		}

		fileDiffs = append(fileDiffs, fileDiff)
	}
	return fileDiffs, nil
}

func (r Repository) executeGitCommand(args ...string) (string, error) {
	return r.commandExecutor.Execute("git", args...)
}

func NewRepository(rootPath string) Repository {
	return newRepository(ShellExecutor{}, rootPath)
}

func newRepository(commandExecutor CommandExecutor, rootPath string) Repository {
	return Repository{
		commandExecutor: commandExecutor,
		rootPath:        rootPath,
	}
}
