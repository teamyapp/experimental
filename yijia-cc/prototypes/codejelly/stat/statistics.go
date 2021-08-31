package stat

import "github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"


func ComputeFileDiffStats(fullFileDiff entity.FullFileDiff) entity.FileDiffStatistics {
	chunks := fullFileDiff.Chunks

	var numOfDeletedLines int
	var numOfAddedLines int
	var fromFileTotalNumOfLines int
	var toFileTotalNumOfLines int

	for _, chunk := range chunks {
		for _, line := range chunk.NumberedLines {
			switch line.Status {
			case entity.LineDeleted:
				numOfDeletedLines++
				fromFileTotalNumOfLines++
			case entity.LineAdded:
				numOfAddedLines++
				toFileTotalNumOfLines++
			default:
				fromFileTotalNumOfLines++
				toFileTotalNumOfLines++
			}

		}
	}

	return entity.FileDiffStatistics{
		NumOfDeletedLines: numOfDeletedLines,
		NumOfAddedLines: numOfAddedLines,
		NumOfLinesChanged: numOfAddedLines + numOfDeletedLines,
		FromFileTotalNumOfLines: fromFileTotalNumOfLines,
		ToFileTotalNumOfLines: toFileTotalNumOfLines,
	}
}

func ComputePullRequestStats(pullRequest entity.PullRequest) entity.FileDiffStatistics {
	panic("implement me")
}
