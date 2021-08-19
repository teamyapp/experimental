package git

import (
	"errors"
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

	return r.parseFileDiffsFromOutput(output)
}

func (r Repository) parseFileDiffsFromOutput(output string) ([]entity.FileDiff, error){
	if len(output) == 0 {
		return nil, errors.New("diff is empty")
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

func(r Repository) GetHunksBetweenBranches(fromBranch string, toBranch string) ([]entity.Hunk, error) {
	output, err := r.executeGitCommand("diff", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	return r.parseHunksFromOutput(output)
}

func (r Repository) parseHunksFromOutput(output string) ([]entity.Hunk, error){
	if len(output) == 0 {
		return nil, errors.New("diff is empty")
	}

	blocks := strings.Split(output, "diff --git")

	hunks := make([]entity.Hunk, 0)

	for _, block := range blocks {
		if len(block) == 0 {
			continue
		}

		hunk, err := newHunkFromLine(block)
		if err != nil {
			return nil, err
		}

		hunks = append(hunks, hunk)
	}

	return hunks, nil
}


func (r Repository) executeGitCommand(args ...string) (string, error) {
	return r.commandExecutor.Execute("git", args...)
}

func NewRepository(rootPath string) Repository {
	return NewRepositoryCustomExecutor(ShellExecutor{}, rootPath)
}

func NewRepositoryCustomExecutor(commandExecutor CommandExecutor, rootPath string) Repository {
	return Repository{
		commandExecutor: commandExecutor,
		rootPath:        rootPath,
	}
}
