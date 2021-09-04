package service

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"strings"
)

func GetChunks(hunks []entity.Hunk, fileContent string) []entity.Chunk{
	if len(hunks) == 0 {
		return nil
	}

	var fromFileLines []string
	if len(fileContent) > 0 {
		fromFileLines = strings.Split(fileContent, "\n")
	}
	numOfLines := len(fromFileLines)
	chunkFromFileStart := 0
	chunkToFileStart := 0
	var chunk entity.Chunk


	chunks := make([]entity.Chunk, 0)

	// TODO: add comments to describe the algorithm with examples

	for _, hunk := range hunks {
		if chunkFromFileStart < numOfLines {
			hunkStart := hunk.FromFileStartLine - 1
			chunk, chunkFromFileStart, chunkToFileStart = getChunk(chunkFromFileStart, hunkStart - 1, chunkToFileStart, fromFileLines)

			if len(chunk.NumberedLines) != 0 {
				chunks = append(chunks, chunk)
			}
		}
		chunk, chunkFromFileStart, chunkToFileStart = hunkToChunk(chunkFromFileStart, chunkToFileStart, hunk)
		chunks = append(chunks, chunk)
	}

	if chunkFromFileStart < numOfLines {
		chunk, _, _ = getChunk(chunkFromFileStart, numOfLines - 1, chunkToFileStart, fromFileLines)
		chunks = append(chunks, chunk)
	}

	return chunks
}

func getChunk(chunkFromFileStart int, chunkFromFileEnd int, chunkToFileStart int, fromFileLines []string) (entity.Chunk, int, int) {
	chunkNumberedLines := make([]entity.NumberedLine, 0)

	fromFileStart := chunkFromFileStart
	fromFileEnd := chunkFromFileEnd
	toFileStart := chunkToFileStart

	for i := fromFileStart; i <= fromFileEnd; i++ {
		numberedLine := entity.NumberedLine {
			Status: entity.LineUnchanged,
			Content: fromFileLines[i],
			FromFileLineNumber: i + 1,
			ToFileLineNumber: i - fromFileStart + toFileStart + 1,
		}

		chunkNumberedLines = append(chunkNumberedLines, numberedLine)
		chunkFromFileStart++
		chunkToFileStart++
	}

	return entity.Chunk{
		NumberedLines: chunkNumberedLines,
		IsHunk: false,
	}, chunkFromFileStart, chunkToFileStart
}

func hunkToChunk(chunkFromFileStart int, chunkToFileStart int, hunk entity.Hunk) (entity.Chunk, int, int) {
	chunkLines := make([]entity.NumberedLine, 0)
	for _, hunkLine := range hunk.Lines {
		status := hunkLine.Status

		var hunkNumberedLine entity.NumberedLine
		switch status {
		case entity.LineUnchanged:
			hunkNumberedLine = entity.NumberedLine{
				Status: status,
				Content: hunkLine.Content,
				FromFileLineNumber: chunkFromFileStart + 1,
				ToFileLineNumber: chunkToFileStart + 1,
			}
			chunkFromFileStart++
			chunkToFileStart++
		case entity.LineDeleted:
			hunkNumberedLine = entity.NumberedLine{
				Status: status,
				Content: hunkLine.Content,
				FromFileLineNumber: chunkFromFileStart + 1,
				ToFileLineNumber: entity.NoLineNumber,
			}
			chunkFromFileStart++
		case entity.LineAdded:
			hunkNumberedLine = entity.NumberedLine{
				Status: status,
				Content: hunkLine.Content,
				FromFileLineNumber: entity.NoLineNumber,
				ToFileLineNumber: chunkToFileStart + 1,
			}
			chunkToFileStart++
		default:
			fmt.Println("error")
		}

		chunkLines = append(chunkLines, hunkNumberedLine)
	}

	return entity.Chunk {
		NumberedLines: chunkLines,
		IsHunk: true,
	}, chunkFromFileStart, chunkToFileStart
}

