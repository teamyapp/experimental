package entity

import (
	"errors"
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

func (c ChangeStatus) String() (string, error) {
	statusName, ok := changeStatusNames[c]
	if !ok {
		return "Invalid", errors.New("invalid status")
	}

	return statusName, nil
}

func NewChangeStatus(statusRune rune) (ChangeStatus, error) {
	status, ok := changeStatusMap[statusRune]

	if !ok {
		return -1, errors.New("invalid status")
	}

	return status, nil
}
