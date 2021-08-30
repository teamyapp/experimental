package entity

import "fmt"

// layer 1: git diff
type Line struct {
	Status  LineStatus
	Content string
}

func (l Line) String() string {
	return fmt.Sprintf("[Line Status=%s, Content=%s]",
		l.Status, l.Content)
}

const NoLineNumber = -1

type NumberedLine struct {
	Status  LineStatus
	Content string
	FromFileLineNumber int
	ToFileLineNumber int
}

func (nl NumberedLine) String() string {
	return fmt.Sprintf("[NumberedLine Status=%s, FromFileLineNumber=%d, ToFileLineNumber=%d, Content=%s]",
		nl.Status, nl.FromFileLineNumber, nl.ToFileLineNumber, nl.Content)
}