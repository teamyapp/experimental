package entity

import "fmt"

//layer 4: calculate statistics
type FileDiffStatistics struct {
	NumOfLinesChanged   int `json:"num_of_lines_changed"`
	NumOfDeletedLines int `json:"num_of_deleted_lines"`
	NumOfAddedLines   int `json:"num_of_added_lines"`
	FromFileTotalNumOfLines int `json:"from_file_total_num_of_lines"`
	ToFileTotalNumOfLines int `json:"to_file_total_num_of_lines"`
}

type FileDiffMetadata struct {
	FileDiffHeader FileDiffHeader `json:"file_diff_header"`
	FileDiffStatistics FileDiffStatistics `json:"file_diff_statistics"`
}

type PullRequestStatistics struct {
	numOfConversations int
	numOfCommits int
	numOfFileChanges int

	totalNumOfDeletedLines int
	totalNumOfAddedLines int
}

func (f FileDiffStatistics) String() string {
	return fmt.Sprintf("[FileDiffStatistics NumOfLinesChanged=%d, NumOfDeletedLines=%d, NumOfAddedLines=%d, FromFileTotalNumOfLines=%d, ToFileTotalNumOfLines=%d]",
		f.NumOfLinesChanged, f.NumOfDeletedLines, f.NumOfAddedLines, f.FromFileTotalNumOfLines, f.ToFileTotalNumOfLines)
}