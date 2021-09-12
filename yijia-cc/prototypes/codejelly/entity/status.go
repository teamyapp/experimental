package entity

import (
	"errors"
	"fmt"
)

type ChangeStatus int

const (
	ChangeAdded ChangeStatus = iota
	ChangeModified
	ChangeDeleted
	ChangeRenamed
)

var changeStatusNames = map[ChangeStatus]string{
	ChangeAdded:    "Added",
	ChangeModified: "Modified",
	ChangeDeleted:  "Deleted",
	ChangeRenamed:  "Renamed",
}

var changeStatusMap = map[rune]ChangeStatus{
	'M': ChangeModified,
	'A': ChangeAdded,
	'D': ChangeDeleted,
	'R': ChangeRenamed,
}

func (c ChangeStatus) String() string {
	statusName, ok := changeStatusNames[c]
	if !ok {
		fmt.Println("change status not found")
		return "unknown"
	}

	return statusName
}

func NewChangeStatus(statusRune rune) (ChangeStatus, error) {
	status, ok := changeStatusMap[statusRune]
	if !ok {
		return -1, errors.New("status not found")
	}

	return status, nil
}

type LineStatus int

const (
	LineUnchanged LineStatus = iota
	LineDeleted
	LineAdded
	LineNothing
)

var lineStatusNames = map[LineStatus]string{
	LineUnchanged:    "Unchanged",
	LineDeleted: "Deleted",
	LineAdded:  "Added",
	LineNothing: "Nothing",
}

func (l LineStatus) String() string {
	statusName, ok := lineStatusNames[l]
	if !ok {
		fmt.Println("line status not found")
		return "unknown"
	}

	return statusName
}