package service

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/vcs"
)

type Diff struct {
	repo vcs.Repository
}

func (d Diff) GetFile(fileDiff entity.FileDiff, fromBranch string) (entity.FullFileDiff, error){
	fromFilePath := fileDiff.FileDiffHeader.FromFilePath
	toFilePath := fileDiff.FileDiffHeader.ToFilePath

	filePath, err := pickFilePath(fromFilePath, toFilePath)
	if err != nil {
		return entity.FullFileDiff{}, err
	}

	fileContent, err := d.repo.GetTextFileContentFromBranch(fromBranch, filePath)
	if err != nil {
		return entity.FullFileDiff{}, err
	}

	chunks := GetChunks(fileDiff.Hunks, fileContent)

	return entity.FullFileDiff {
		FileDiffHeader: fileDiff.FileDiffHeader,
		Chunks: chunks,
	}, nil
}

func NewDiff(repo vcs.Repository) Diff {
	return Diff{
		repo: repo,
	}
}
