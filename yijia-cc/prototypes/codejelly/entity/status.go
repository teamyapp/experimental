package entity

import (
	"errors"
	"log"
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
		log.Print("change status not found")
		return ""
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
