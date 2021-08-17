package git

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"os/exec"
	"strings"
)

type Repository struct {
	rootPath string
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

func (r Repository) executeGitCommand(args ...string) (string, error){
	args = append([]string{"-C", r.rootPath}, args... )
	out, err := exec.Command("git", args...).Output()
	//cmd.Stderr = os.Stdout
	return string(out), err
}

func NewRepository(rootPath string) Repository {
	return Repository{
		rootPath: rootPath,
	}
}

