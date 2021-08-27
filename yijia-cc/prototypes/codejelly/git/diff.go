package git

import (
	"errors"
	"fmt"
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


//var filePathPattern = regexp.MustCompile("(/)+[a-zA-Z0-9\\\\-_/ ]*(.)")
const (
	// e.g. "@@ -17,2 +11,3 @@"
	lineChangeStatPrefix string = "@@"
	// e.g. "index 0000000..edc67e8 "
	indexPrefix string = "index"
	// e.g. "similarity index 58%"
	similarityPrefix string = "similarity"
	// e.g. "rename from discussion/src/main/java/info/grouplive/discussion/model/User.java"
	renameFromPrefix string = "rename from"
	// e.g. "rename to discussion/src/main/java/info/grouplive/discussion/model/UserModel.java"
	renameToPrefix string = "rename to"
	// e.g. "--- /dev/null"
	oldFilePathPrefix string = "--- "
	// e.g. "+++ b/calendar/repo/reservation.go"
	newFilePathPrefix string = "+++ "
	// e.g. "\ No newline at end of file"
	noNewLinePrefix string = "\\ No newline at end of file"

	// e.g. "Binary files /dev/null and b/discussion/dependencies/proto-1.0-SNAPSHOT.jar differ"
	binaryFilePrefix string = "Binary"

	// unchanged line starts with an empty space
	// e.g. "     public UserDetails loadUserByUsername(String username) {"
	noChangeLinePrefix string = " "
	// e.g. "+        UserModel user = userOptional"
	addedLinePrefix string = "+"
	// e.g. "-        User user = userOptional"
	deletedLinePrefix string = "-"
)

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

	var hunkHeader entity.HunkHeader
	var err error

	for index, line := range lines {
		if index == 0 {
			continue
		}

		// index info is not needed, hence it is skipped
		if strings.HasPrefix(line, noNewLinePrefix) ||
			strings.HasPrefix(line, indexPrefix) {
			continue
		}

		if strings.HasPrefix(line, lineChangeStatPrefix) {
			hunks = tryAddPrevHunk(hunkCounter, hunks, hunkHeader, hunkLines)

			hunkHeader, err = parseLineIntoStatString(line)
			if err != nil {
				return nil, err
			}

			hunkCounter += 1
			hunkLines = make([]entity.Line, 0)
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineHunkHeader,
				Content: line,
			})

		} else if strings.HasPrefix(line, noChangeLinePrefix){
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineUnchanged,
				Content: line,
			})
		} else if strings.HasPrefix(line, deletedLinePrefix){
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineDeleted,
				Content: line,
			})
		} else if strings.HasPrefix(line, addedLinePrefix){
			hunkLines = append(hunkLines, entity.Line{
				Status: entity.LineAdded,
				Content: line,
			})
		}
	}

	hunks = tryAddPrevHunk(hunkCounter, hunks, hunkHeader, hunkLines)

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
			continue
		}

		if strings.HasPrefix(line, similarityPrefix) {
			similarity, err = parseSimilarityFromLine(line)
			if err != nil {
				return entity.FileDiffHeader{}, err
			}
		} else if strings.HasPrefix(line, renameFromPrefix) || strings.HasPrefix(line, oldFilePathPrefix){
			fromFilePath = parseFileNameFromLine(line)
		} else if strings.HasPrefix(line, renameToPrefix) || strings.HasPrefix(line, newFilePathPrefix){
			toFilePath = parseFileNameFromLine(line)
		} else if strings.HasPrefix(line, binaryFilePrefix) {
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

func tryAddPrevHunk(
	hunkCounter int,
	hunks []entity.Hunk,
	hunkHeader entity.HunkHeader,
	hunkLines []entity.Line) []entity.Hunk{
	if hunkCounter == 0 {
		return hunks
	}

	return append(hunks, entity.Hunk {
		HunkHeader: hunkHeader,
		Lines: append([]entity.Line(nil), hunkLines...),
	})
}

