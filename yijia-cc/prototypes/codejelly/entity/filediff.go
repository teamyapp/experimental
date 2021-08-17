package entity

import "fmt"

type FileDiff struct {
	Status       ChangeStatus
	FromFilePath string
	ToFilePath   string
	Similarity   int
}

func (f FileDiff) String() string{
	return fmt.Sprintf("[FileDiff Status=%s, FromFilePath=%s, ToFilePath=%s, Similarity=%d",
		f.Status, f.FromFilePath, f.ToFilePath, f.Similarity)
}

