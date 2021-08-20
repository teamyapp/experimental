package entity

import "fmt"

type FileDiffHeader struct {
	Id int
	Status       ChangeStatus
	FromFilePath string
	ToFilePath   string
	Similarity   int
}

func (f FileDiffHeader) String() string {
	return fmt.Sprintf("[FileDiffHeader Status=%s, FromFilePath=%s, ToFilePath=%s, Similarity=%d",
		f.Status, f.FromFilePath, f.ToFilePath, f.Similarity)
}

type FileDiff struct {
	FileDiffHeaderId int
	Hunks []Hunk
}