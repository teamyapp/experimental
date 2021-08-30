package vcs

import "github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"

type Repository interface {
	GetFileDiffHeadersBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiffHeader, error)
	GetFileDiffsBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiff, error)
	GetTextFileContentFromBranch(branch string, filepath string) (string, error)
}
