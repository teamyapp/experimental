package git

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func newFileDiffHeaderFromLine(line string, index int) (entity.FileDiffHeader, error) {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		log.Println("line must have at least 2 parts")
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
		Id:           index,
		Status:       status,
		FromFilePath: parts[1],
		ToFilePath:   toFilePath,
		Similarity:   similarity,
	}, nil
}

