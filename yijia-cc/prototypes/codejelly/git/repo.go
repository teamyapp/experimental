package git

import (
	"errors"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/vcs"
)

type Git struct {
	repoRootPath        string
	commandExecutor CommandExecutor
}

var _ vcs.VersionControlSystem = (*Git)(nil)

func (g Git) GetFileDiffHeadersBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiffHeader, error) {
	output, err := g.executeGitCommand("diff", "--name-status", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	return g.parseFileDiffHeadersFromOutput(output)
}

func (g Git) parseFileDiffHeadersFromOutput(output string) ([]entity.FileDiffHeader, error){
	if len(output) == 0 {
		return nil, errors.New("diff is empty")
	}

	lines := strings.Split(output, "\n")

	fileDiffs := make([]entity.FileDiffHeader, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fileDiff, err := newFileDiffHeaderFromLine(line, len(fileDiffs))
		if err != nil {
			return nil, err
		}

		fileDiffs = append(fileDiffs, fileDiff)
	}
	return fileDiffs, nil
}

func(g Git) GetFileDiffsBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiff, error) {
	output, err := g.executeGitCommand("diff", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	return g.parseFileDiffsFromOutput(output)
}

func (g Git) parseFileDiffsFromOutput(output string) ([]entity.FileDiff, error){
	if len(output) == 0 {
		return nil, errors.New("diff is empty")
	}

	// each file has a code of block, containing filename, file change stats, a slice of hunks for the file, etc.
	blocks := strings.Split(output, "diff --git")

	fileDiffs := make([]entity.FileDiff, 0)

	for _, block := range blocks {
		if len(block) == 0 {
			continue
		}

		fileDiff, err := newFileDiffFromBlock(block, len(fileDiffs))
		if err != nil {
			return nil, err
		}

		fileDiffs = append(fileDiffs, fileDiff)
	}

	return fileDiffs, nil
}


func (g Git) executeGitCommand(args ...string) (string, error) {
	return g.commandExecutor.Execute("git", args...)
}

func NewGit(repoRootPath string) Git {
	return NewGitCustomExecutor(ShellExecutor{}, repoRootPath)
}

func NewGitCustomExecutor(commandExecutor CommandExecutor, repoRootPath string) Git {
	return Git{
		commandExecutor: commandExecutor,
		repoRootPath:        repoRootPath,
	}
}
