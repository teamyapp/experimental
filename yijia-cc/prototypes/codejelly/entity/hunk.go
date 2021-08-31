package entity

import "fmt"

// layer 2: unorganized hunks
type Hunk struct {
	//index string
	HunkHeader
	Lines []Line // - line 16-37, + line 16-33, unchanged
}

type HunkHeader struct {
	FromFileStartLine  int
	FromFileNumOfLines int
	ToFileStartLine    int
	ToFileNumOfLines   int
	HeaderLine string
}

func (h HunkHeader) String() string {
	return fmt.Sprintf("[HunkHeader FromFileStartLine=%d, FromFileNumOfLines=%d, ToFileStartLine=%d, ToFileNumOfLines=%d",
		h.FromFileStartLine, h.FromFileNumOfLines, h.ToFileStartLine, h.ToFileNumOfLines)
}
