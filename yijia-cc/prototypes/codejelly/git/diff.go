package git

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func newFileDiffFromLine(line string) (entity.FileDiff, error) {
	// throw error:
	// 		invalid status
	// 		cannot parse similarity
	// skip parsing:
	// 		empty line
	// 		num of line parts < 2
	parts := strings.Fields(line)
	if len(parts) < 2 {
		log.Println("line must have at least 2 parts")
		return entity.FileDiff{}, nil
	}

	status, err := entity.NewChangeStatus(rune(parts[0][0]))
	if err != nil {
		return entity.FileDiff{}, errors.New("invalid status")
	}

	var toFilePath string
	var similarity int
	switch status {
	case entity.ChangeRenamed:
		similarity, err = strconv.Atoi(parts[0][1:])
		if err != nil {
			return entity.FileDiff{}, err
		}
		toFilePath = parts[2]
	case entity.ChangeAdded, entity.ChangeDeleted, entity.ChangeModified:
		toFilePath = parts[1]
	default:
		return entity.FileDiff{}, errors.New("invalid status")
	}

	return entity.FileDiff{
		Status:       status,
		FromFilePath: parts[1],
		ToFilePath:   toFilePath,
		Similarity:   similarity,
	}, nil
}
