package service

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/vcs"
	"strings"
)

type CodeReview struct {
	repo vcs.Repository
}

func (c CodeReview) GetFile(fileDiff entity.FileDiff, fromBranch string) (entity.FullFileDiff, error){
	output, err := c.repo.GetTextFileContentFromBranch(fromBranch, fileDiff.FileDiffHeader.FromFilePath)
	if err != nil {
		return entity.FullFileDiff{}, err
	}

	lines := strings.Split(output, "\n")
	chunks := getChunks(fileDiff.Hunks, lines)

	return entity.FullFileDiff {
		FileDiffHeader: fileDiff.FileDiffHeader,
		Chunks: chunks,
	}, nil
}

func NewCodeReview(repo vcs.Repository) CodeReview {
	return CodeReview{
		repo: repo,
	}
}

