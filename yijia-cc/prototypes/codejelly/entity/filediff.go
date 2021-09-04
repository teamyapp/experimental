package entity

import "fmt"

const GitNullFile = "/dev/null"

type FileDiffHeader struct {
	Status       ChangeStatus `json:"status"`
	FromFilePath string `json:"from_file_path"`
	ToFilePath   string `json:"to_file_path"`
	Similarity   int `json:"similarity"`
}

func (f FileDiffHeader) String() string {
	return fmt.Sprintf("[FileDiffHeader Status=%s, FromFilePath=%s, ToFilePath=%s, Similarity=%d",
		f.Status, f.FromFilePath, f.ToFilePath, f.Similarity)
}

type FileDiff struct {
	FileDiffHeader
	Hunks []Hunk
	HasNoNewLineSymbol bool
}
