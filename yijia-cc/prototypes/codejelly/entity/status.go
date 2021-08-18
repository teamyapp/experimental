package entity

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
	return changeStatusNames[c]
}

func NewChangeStatus(statusRune rune) ChangeStatus {
	return changeStatusMap[statusRune]
}
