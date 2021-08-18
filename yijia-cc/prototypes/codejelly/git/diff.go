package git

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"strconv"
	"strings"
)

func newFileDiffFromLine(line string) (entity.FileDiff, error) {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return entity.FileDiff{}, fmt.Errorf("line must have at least 2 parts")
	}

	status := entity.NewChangeStatus(rune(parts[0][0]))
	var toFilePath string
	var similarity int
	var err error
	switch status {
	case entity.ChangeRenamed:
		similarity, err = strconv.Atoi(parts[0][1:])
		if err != nil {
			return entity.FileDiff{}, err
		}
		toFilePath = parts[2]
	default:
		toFilePath = parts[1]
	}

	return entity.FileDiff{
		Status:       status,
		FromFilePath: parts[1],
		ToFilePath:   toFilePath,
		Similarity:   similarity,
	}, nil

}
