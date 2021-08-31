package view

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/stat"
)

func RenderUnifiedView(fullFileDiff entity.FullFileDiff) entity.UnifiedView {
	fileDiffStats := stat.ComputeFileDiffStats(fullFileDiff)

	indices := make([]int, 0)
	for index, chunk := range fullFileDiff.Chunks {
		if chunk.IsHunk {
			indices = append(indices, index)
		}
	}

	return entity.UnifiedView{
		FileDiffMetadata: entity.FileDiffMetadata{
			FileDiffHeader:     fullFileDiff.FileDiffHeader,
			FileDiffStatistics: fileDiffStats,
		},
		AllChunks:   fullFileDiff.Chunks,
		HunkIndices: indices,
	}
}

func RenderSplitView(fullFileDiff entity.FullFileDiff) entity.SplitView {
	fileDiffStats := stat.ComputeFileDiffStats(fullFileDiff)

	indices := make([]int, 0)
	chunkPairs := make([]entity.ChunkPair, 0)

	for index, chunk := range fullFileDiff.Chunks {
		if chunk.IsHunk {
			indices = append(indices, index)
		}
		chunkPair := getChunkPair(chunk)
		chunkPairs = append(chunkPairs, chunkPair)
	}

	return entity.SplitView{
		FileDiffMetadata: entity.FileDiffMetadata{
			FileDiffHeader:     fullFileDiff.FileDiffHeader,
			FileDiffStatistics: fileDiffStats,
		},
		AllChunkPairs:   chunkPairs,
		HunkPairIndices: indices,
	}
}

func getChunkPair(chunk entity.Chunk) entity.ChunkPair {
	if !chunk.IsHunk {
		return renumberedChunkPair(chunk)
	}

	return hunkToChunkPair(chunk)
}

func renumberedChunkPair(chunk entity.Chunk) entity.ChunkPair {
	fromFileNumberedLines := make([]entity.NumberedLine, 0)
	toFileNumberedLines := make([]entity.NumberedLine, 0)

	for _, numberedLine := range chunk.NumberedLines {
		fromFileNewNumberedLine := createNumberedLine(
			entity.LineUnchanged,
			numberedLine.Content,
			numberedLine.FromFileLineNumber,
			entity.NoLineNumber,
		)
		toFileNewNumberedLine := createNumberedLine(
			entity.LineUnchanged,
			numberedLine.Content,
			entity.NoLineNumber,
			numberedLine.ToFileLineNumber,
		)
		fromFileNumberedLines = append(fromFileNumberedLines, fromFileNewNumberedLine)
		toFileNumberedLines = append(toFileNumberedLines, toFileNewNumberedLine)
	}

	return entity.ChunkPair{
		FromFileChunk: entity.Chunk{
			NumberedLines: fromFileNumberedLines,
			IsHunk: chunk.IsHunk,
		},
		ToFileChunk: entity.Chunk{
			NumberedLines: toFileNumberedLines,
			IsHunk: chunk.IsHunk,
		},
	}
}

func hunkToChunkPair(chunk entity.Chunk) entity.ChunkPair {

	fromFileNumberedLines := make([]entity.NumberedLine, 0)
	toFileNumberedLines := make([]entity.NumberedLine, 0)

	deletedLineCounter := 0
	addedLineCounter := 0

	for _, numberedLine := range chunk.NumberedLines {
		switch numberedLine.Status {
		case entity.LineUnchanged:
			fromFileNewNumberedLine := createNumberedLine(
				entity.LineUnchanged,
				numberedLine.Content,
				numberedLine.FromFileLineNumber,
				entity.NoLineNumber,
			)
			toFileNewNumberedLine := createNumberedLine(
				entity.LineUnchanged,
				numberedLine.Content,
				entity.NoLineNumber,
				numberedLine.ToFileLineNumber,
			)

			if deletedLineCounter < addedLineCounter {
				fromFileNumberedLines = appendNothingLine(fromFileNumberedLines, addedLineCounter - deletedLineCounter)
			} else if deletedLineCounter > addedLineCounter {
				toFileNumberedLines = appendNothingLine(toFileNumberedLines, deletedLineCounter - addedLineCounter)
			}

			fromFileNumberedLines = append(fromFileNumberedLines, fromFileNewNumberedLine)
			toFileNumberedLines = append(toFileNumberedLines, toFileNewNumberedLine)
			deletedLineCounter = 0
			addedLineCounter = 0
		case entity.LineDeleted:
			fromFileNewNumberedLine := createNumberedLine(
				entity.LineDeleted,
				numberedLine.Content,
				numberedLine.FromFileLineNumber,
				entity.NoLineNumber,
			)
			fromFileNumberedLines = append(fromFileNumberedLines, fromFileNewNumberedLine)
			deletedLineCounter++
		case entity.LineAdded:
			toFileNumberedLine := createNumberedLine(
				entity.LineAdded,
				numberedLine.Content,
				entity.NoLineNumber,
				numberedLine.ToFileLineNumber,
			)
			toFileNumberedLines = append(toFileNumberedLines, toFileNumberedLine)
			addedLineCounter++
		default:
		}
	}
	return entity.ChunkPair{
		FromFileChunk: entity.Chunk{
			NumberedLines: fromFileNumberedLines,
			IsHunk: true,
		},
		ToFileChunk: entity.Chunk{
			NumberedLines: toFileNumberedLines,
			IsHunk: true,
		},
	}
}

func createNumberedLine(
	status entity.LineStatus,
	content string,
	fromFileLineNumber int,
	toFileLineNumber int) entity.NumberedLine {
	return entity.NumberedLine{
		Status:             status,
		Content:            content,
		FromFileLineNumber: fromFileLineNumber,
		ToFileLineNumber:   toFileLineNumber,
	}
}

func appendNothingLine(lines []entity.NumberedLine, num int) []entity.NumberedLine{
	for i := 0; i < num; i++ {
		lines = append(lines, entity.NumberedLine{
			Status: entity.LineNothing,
			Content: "",
			FromFileLineNumber: entity.NoLineNumber,
			ToFileLineNumber: entity.NoLineNumber,
		})
	}
	return lines
}
