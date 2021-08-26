package git

import (
	"errors"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"strconv"
	"strings"
)

func parseFileNameFromLine (line string) string{
	parts := strings.Fields(line)

	return trimFileName(parts[len(parts) - 1])
}

func trimFileName (fileName string) string {
	if fileName[:2] == "a/" {
		fileName = strings.TrimPrefix(fileName, "a/")
	}
	if fileName[:2] == "b/" {
		fileName = strings.TrimPrefix(fileName, "b/")
	}
	return fileName
}

func parseSimilarityFromLine(line string) (int, error) {
	parts := strings.Fields(line)
	numberString := strings.TrimRight(parts[len(parts) - 1], "%")
	num, err := strconv.Atoi(numberString)

	if err != nil {
		return -1, err
	}
	return num, nil
}

func parseLineIntoStatString (line string) (int, int, int, int, error) {
	// line example: "@@ -17,6 +16,7 @@ type Calendar struct {"
	// return: 17, 6, 16, 7

	// stats example: ["@@", "-17,6", "+16,7", "@@", ....]
	stats := strings.Fields(line)

	// fromFileStatString example: "-17,6"
	fromFileStatString := stats[1]
	// toFileStatString example: "+16,7"
	toFileStatString := stats[2]

	/*
	var err error

	f := func(statString string) bool {
		_, _, err := parseLineStatsFromStatString(statString)
		return err == nil
	}

	if !f(fromFileStatString) && !f(toFileStatString) {
		return -1, -1, -1, -1, err
	}
	*/

	fromFileStartLine, fromFileNumOfLines, err := parseLineStatsFromStatString(fromFileStatString)
	if err != nil {
		return -1, -1, -1, -1, err
	}
	toFileStartLine, toFileNumOfLines, err := parseLineStatsFromStatString(toFileStatString)
	if err != nil {
		return -1, -1, -1, -1, err
	}

	return fromFileStartLine, fromFileNumOfLines, toFileStartLine, toFileNumOfLines, nil
}

func parseLineStatsFromStatString (statString string) (int, int, error) {

	// statString sample: "-17,6"
	// return: 17, 6
	//fmt.Println(statString)

	// To handle "+1" case
	if statString == "+1" || statString == "-1"{
		return 1, 1, nil
	}

	statParts := strings.Split(statString, ",")
	if len(statParts) != 2 {
		return -1, -1, errors.New("invalid stats string")
	}

	startLineString := statParts[0]
	numOfLinesString := statParts[1]

	var err error

	f := func(s string) bool {
		_, err := getPositiveIntegerFromString(s)
		return err == nil
	}

	if !f(startLineString) && !f(numOfLinesString) {
		return -1, -1, err
	}

	startLine, _ := getPositiveIntegerFromString(startLineString)
	numOfLines, _ := getPositiveIntegerFromString(numOfLinesString)

	return startLine, numOfLines, nil
}


func getFileStatusFromFileName(fromFilePath string, toFilePath string) entity.ChangeStatus{
	var status entity.ChangeStatus

	if fromFilePath == toFilePath && fromFilePath != "/dev/null"{
		status = entity.ChangeModified
		return status
	}

	if fromFilePath == "/dev/null" && toFilePath != "/dev/null"{
		status = entity.ChangeAdded
		return status
	}

	if toFilePath == "/dev/null" && fromFilePath != "/dev/null" {
		status = entity.ChangeDeleted
		return status
	}

	return entity.ChangeRenamed
}

func getPositiveIntegerFromString (s string) (int, error) {
	// s example: "-16",
	// return 16
	output, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}

	return absolute(output), nil
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

