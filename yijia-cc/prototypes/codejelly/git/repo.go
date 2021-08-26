package git

import (
	"errors"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/vcs"
)

type Repository struct {
	repoRootPath        string
	commandExecutor CommandExecutor
}

var _ vcs.Repository = (*Repository)(nil)

func (r Repository) GetFileDiffHeadersBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiffHeader, error) {
	output, err := r.executeGitCommand("diff", "--name-status", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	return r.parseFileDiffHeadersFromOutput(output)
}


func (r Repository) parseFileDiffHeadersFromOutput(output string) ([]entity.FileDiffHeader, error){
	if len(output) == 0 {
		return nil, errors.New("diff is empty")
	}

	lines := strings.Split(output, "\n")

	fileDiffs := make([]entity.FileDiffHeader, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fileDiff, err := newFileDiffHeaderFromLine(line)
		if err != nil {
			return nil, err
		}

		fileDiffs = append(fileDiffs, fileDiff)
	}
	return fileDiffs, nil
}


func(r Repository) GetFileDiffsBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiff, error) {
	output, err := r.executeGitCommand("diff", fromBranch, toBranch)
	if err != nil {
		return nil, err
	}

	return r.parseFileDiffsFromOutput(output)
}


func (r Repository) parseFileDiffsFromOutput(output string) ([]entity.FileDiff, error){
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

		fileDiff, err := newFileDiffFromBlock(block)
		if err != nil {
			return nil, err
		}

		fileDiffs = append(fileDiffs, fileDiff)
	}

	return fileDiffs, nil
}

func (r Repository) executeGitCommand(args ...string) (string, error) {
	args = append([]string{"-C", r.repoRootPath}, args...)
	return r.commandExecutor.Execute("git", args...)
}

func NewRepo(repoRootPath string) Repository {
	shellExecutor := ShellExecutor{}
	return NewRepositoryDeps(shellExecutor, repoRootPath)
}

func NewRepositoryDeps(commandExecutor CommandExecutor, repoRootPath string) Repository {
	return Repository{
		commandExecutor: commandExecutor,
		repoRootPath:    repoRootPath,
	}
}


/*Special Cases for parsing
FileDiffHeader:
- TODO: Binary files:
	- Binary files /dev/null and b/discussion/dependencies/proto-1.0-SNAPSHOT.jar differ
	- Binary file change file extension to txt
	- Binary file change content but keep name as same
	- Binary file deleted
- TODO: image file
- TODO: audio file

StatsString:
- TODO: StatsString have +1 only, equivalent to +1,1

*/