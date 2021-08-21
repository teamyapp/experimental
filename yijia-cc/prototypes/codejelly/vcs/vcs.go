package vcs

import "github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"

type VersionControlSystem interface {
	GetFileDiffHeadersBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiffHeader, error)
	GetFileDiffsBetweenBranches(fromBranch string, toBranch string) ([]entity.FileDiff, error)
}
