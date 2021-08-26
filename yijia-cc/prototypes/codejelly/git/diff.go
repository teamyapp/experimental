package git

import (
	"errors"
	"fmt"
	"regexp"

	"strconv"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func newFileDiffHeaderFromLine(line string) (entity.FileDiffHeader, error) {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		fmt.Println("line must have at least 2 parts")
		return entity.FileDiffHeader{}, nil
	}

	status, err := entity.NewChangeStatus(rune(parts[0][0]))
	if err != nil {
		return entity.FileDiffHeader{}, err
	}

	var toFilePath string
	var similarity int
	switch status {
	case entity.ChangeRenamed:
		similarity, err = strconv.Atoi(parts[0][1:])
		if err != nil {
			return entity.FileDiffHeader{}, err
		}
		toFilePath = parts[2]
	case entity.ChangeAdded, entity.ChangeDeleted, entity.ChangeModified:
		toFilePath = parts[1]
	default:
		return entity.FileDiffHeader{}, errors.New("invalid status")
	}

	return entity.FileDiffHeader{
		Status:       status,
		FromFilePath: parts[1],
		ToFilePath:   toFilePath,
		Similarity:   similarity,
	}, nil
}

// TODO: add example
//var filePathPattern = regexp.MustCompile("(/)+[a-zA-Z0-9\\\\-_/ ]*(.)")
var lineChangeStatPattern = regexp.MustCompile("^@@*")
var indexPattern = regexp.MustCompile("^index*")
var similarityPattern = regexp.MustCompile("^similarity")
var renameFromPattern = regexp.MustCompile("^rename from")
var renameToPattern = regexp.MustCompile("^rename to")
var noNewLinePattern = regexp.MustCompile("\\ No newline at end of file")
var noChangeLinePattern = regexp.MustCompile("^ ")
var addedLinePattern = regexp.MustCompile("^+")
var deletedLinePattern = regexp.MustCompile("^-")
var binaryFilePattern = regexp.MustCompile("^Binary")

func newHunkFromBlock(block string) ([]entity.Hunk, error) {
	if len(block) == 0 {
		return nil, nil
	}

	block = strings.TrimSpace(block)
	lines := strings.Split(block, "\n")

	// increment hunkCounter when line match lineChangeStatPattern
	hunkCounter := 0
	hunks := make([]entity.Hunk, 0)
	hunkLines := make([]entity.Line, 0)

	var fromFileStartLine int
	var fromFileNumOfLines int
	var toFileStartLine    int
	var toFileNumOfLines   int
	var err error

	for index, line := range lines {
		if index == 0 {
			continue
		}

		if noNewLinePattern.MatchString(line) || indexPattern.MatchString(line){
			continue
		}

		if lineChangeStatPattern.MatchString(line) {
			if hunkCounter != 0 {
				hunks = append(hunks, entity.Hunk {
					FromFileStartLine: fromFileStartLine,
					FromFileNumOfLines: fromFileNumOfLines,
					ToFileStartLine: toFileStartLine,
					ToFileNumOfLines: toFileNumOfLines,

					Lines: append([]entity.Line(nil), hunkLines...),
				})
			}

			fromFileStartLine, fromFileNumOfLines, toFileStartLine, toFileNumOfLines, err = parseLineIntoStatString(line)
			if err != nil {
				return nil, err
			}

			hunkCounter += 1
			hunkLines = make([]entity.Line, 0)
		} else if noChangeLinePattern.MatchString(line){
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineUnchanged,
				Content: line,
			})
		} else if deletedLinePattern.MatchString(line) {
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineDeleted,
				Content: line,
			})
		} else if addedLinePattern.MatchString(line) {
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineAdded,
				Content: line,
			})
		}
	}

	if hunkCounter != 0 {
		hunks = append(hunks, entity.Hunk {
			FromFileStartLine: fromFileStartLine,
			FromFileNumOfLines: fromFileNumOfLines,
			ToFileStartLine: toFileStartLine,
			ToFileNumOfLines: toFileNumOfLines,

			Lines: append([]entity.Line(nil), hunkLines...),
		})
	}

	return hunks, nil
}

func newFileDiffHeaderFromBlock(block string) (entity.FileDiffHeader, error) {
	if len(block) == 0 {
		return entity.FileDiffHeader{}, errors.New("invalid git diff hunk")
	}

	contents := strings.Split(block, "@@")
	content := contents[0]
	lines := strings.Split(content, "\n")

	var status entity.ChangeStatus
	var fromFilePath string
	var toFilePath    string
	var similarity   int
	var err error

	for _, line := range lines {
		if len(line) <= 4 {
			//fmt.Println("line is not valid")
			continue
		}

		if similarityPattern.MatchString(line) {
			similarity, err = parseSimilarityFromLine(line)
			if err != nil {
				return entity.FileDiffHeader{}, err
			}
		} else if  renameFromPattern.MatchString(line) || line[:4] == "--- "{
			fromFilePath = parseFileNameFromLine(line)
		} else if renameToPattern.MatchString(line) || line[:4] == "+++ " {
			toFilePath = parseFileNameFromLine(line)
		} else if binaryFilePattern.MatchString(line) {
			lineParts := strings.Fields(line)
			fromFilePath = trimFileName(lineParts[2])
			toFilePath = trimFileName(lineParts[4])
		}
	}

	status = getFileStatusFromFileName(fromFilePath, toFilePath)

	if len(fromFilePath) == 0 || len(toFilePath) == 0 {
		return entity.FileDiffHeader{}, errors.New("invalid git diff hunk")
	}

	return entity.FileDiffHeader{
		Status: status,
		FromFilePath: fromFilePath,
		ToFilePath: toFilePath,
		Similarity: similarity,
	}, nil
}

func newFileDiffFromBlock(block string) (entity.FileDiff, error) {
	if len(block) == 0 {
		return entity.FileDiff{}, errors.New("file does not have diff")
	}

	fileDiffHeader, err := newFileDiffHeaderFromBlock(block)

	hunks, err := newHunkFromBlock(block)
	if err != nil {
		return entity.FileDiff{}, err
	}

	return entity.FileDiff{
		FileDiffHeader: fileDiffHeader,
		Hunks:            hunks,
	}, nil

}

