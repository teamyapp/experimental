package service

import (
	"errors"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/vcs"
)

type CodeReview struct {
	repo vcs.Repository
}

func (c CodeReview) GetFile(fileDiff entity.FileDiff, fromBranch string) (entity.FullFileDiff, error){
	fromFilePath := fileDiff.FileDiffHeader.FromFilePath
	toFilePath := fileDiff.FileDiffHeader.ToFilePath

	filePath, err := pickFilePath(fromFilePath, toFilePath)
	if err != nil {
		return entity.FullFileDiff{}, err
	}

	fileContent, err := c.repo.GetTextFileContentFromBranch(fromBranch, filePath)
	if err != nil {
		return entity.FullFileDiff{}, err
	}

	chunks := GetChunks(fileDiff.Hunks, fileContent)

	return entity.FullFileDiff {
		FileDiffHeader: fileDiff.FileDiffHeader,
		Chunks: chunks,
	}, nil
}

func pickFilePath (fromFilePath string, toFilePath string) (string, error) {
	if len(fromFilePath) == 0 && len(toFilePath) == 0 {
		return "", errors.New("invalid fileDiff")
	}

	if len(fromFilePath) == 0 {
		return toFilePath, nil
	}

	if len(toFilePath) == 0 {
		return toFilePath, nil
	}

	return toFilePath, nil
}

func NewCodeReview(repo vcs.Repository) CodeReview {
	return CodeReview{
		repo: repo,
	}
}

