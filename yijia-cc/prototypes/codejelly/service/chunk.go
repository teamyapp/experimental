package service

import "github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"

func getChunks(hunks []entity.Hunk, fromFileLines []string) []entity.Chunk{
	if len(hunks) == 0 || len(fromFileLines) == 0 {
		return nil
	}

	chunks := make([]entity.Chunk, 0)
	numOfLines := len(fromFileLines)
	// TODO: add comments to describe the algorithm with examples
	chunkStart := 0
	for _, hunk := range hunks {
		if chunkStart < numOfLines {
			curHunkStart := hunk.FromFileStartLine
			curHunkEnd := curHunkStart + hunk.FromFileNumOfLines - 1
			chunk := getChunk(chunkStart, curHunkStart - 1, fromFileLines)

			if chunk.Lines != nil {
				chunks = append(chunks, chunk)
			}

			chunks = append(chunks, entity.Chunk {
				Lines: hunk.Lines,
				IsHunk: true,
			})

			chunkStart = curHunkEnd + 1
		}
	}

	if chunkStart < numOfLines {
		chunk := getChunk(chunkStart, numOfLines - 1, fromFileLines)
		chunks = append(chunks, chunk)
	}

	return chunks
}

func getChunk(chunkStart int, chunkEnd int, fromFileLines []string) entity.Chunk {
	chunkLines := make([]entity.Line, 0)

	for i := chunkStart; i <= chunkEnd; i++ {
		outputLine := entity.Line {
			Status: entity.LineUnchanged,
			Content: fromFileLines[i],
		}
		chunkLines = append(chunkLines, outputLine)
	}

	return entity.Chunk{
		Lines: chunkLines,
		IsHunk: false,
	}
}
