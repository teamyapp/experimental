package entity

import "time"

//  1. Pull Request Info
type PullRequest struct {
	id          int
	title       string
	description string
	author      User
	reviewers   []Reviewer

	repo             string
	sourceBranch     string
	targetBranch     string
	parentId         *int // -1 root PR/ base PR
	openedAt         time.Time
	availableActions []PullRequestAction

	status   PullRequestStatus
	comments []Comment
	//commits  []Commit

	//aggregatedFileChanges
}

// type Policy struct {}

type PullRequestStatus int

const (
	prOpened PullRequestStatus = iota
	//unresolved
	prApproved
	prMerged
	prClosed
)

type PullRequestAction int

const (
	actionApprove PullRequestAction = iota
	actionClose
	actionMerge
	actionSquashAndMerge
	actionRebaseAndMerge
)

type Reviewer struct {
	user   User
	status ReviewStatus
}

type ReviewStatus int

const (
	rsNotViewed ReviewStatus = iota
	rsUnresolvedComment
	rsApproved
)

type User struct {
	id   int
	name string
}

// 2. Code Comment
type Thread struct {
	id       int
	comments []Comment
	// mark unresolved when new comments added
	isResolved bool
	selection  Selection
}

type Comment struct {
	id         int
	reviewerId int
	// TODO:
	// 		1) allow emojis
	// 		2) allow tagging others
	// 		3) quote code
	// 		4) offer code suggestions
	content   string
	createdAt time.Time
}

type Selection struct {
	startLine   int
	startColumn int
	endLine     int
	endColumn   int
}

// 3. Code Diff
/*
layer 1: git diff
layer 2: unorganized hunks
layer 3: group unchanged chunks and hunks into file change pair
layer 4: feed data for split view and unified view
layer 5: render UI at frontend
*/

// layer 1: git diff
type Line struct {
	Status  LineStatus
	Content string
}

type LineStatus int

const (
	LineUnchanged LineStatus = iota
	LineDeleted
	LineAdded
	LineNothing
)

//type Commit struct {
//	id           int
//	message      string
//	createdAt    time.Time
//	changedFiles []File
//	author       User
//}

//type File struct {
//	id                int
//	filePath          string
//	commitId          int
//	status            ChangeStatus
//	numOfChangedLines int
//	comments          []Comment
//}

// layer 2: unorganized hunks
type Hunk struct {
	//index string
	FromFileStartLine  int
	FromFileNumOfLines int
	ToFileStartLine    int
	ToFileNumOfLines   int

	Lines []Line // - line 16-37, + line 16-33, unchanged
}

//layer 3: group unchanged chunks and hunks into file change pair
type DiffLine struct {
	Line
	oldLineNumber int
	newLineNumber int
}

type DiffFilePaths struct {
	oldFilePath string
	newFilePath string
	isRenamed   bool
}

type Chunk struct {
	lines  []DiffLine
	isHunk bool
}

type Diff struct {
	DiffFilePaths
	chunks []Chunk
}

//layer 4: feed data for split view and unified view
type DiffStatistics struct {
	totalNumOfLines   int
	numOfDeletedLines int
	numOfAddedLines   int
}

type DiffMetadata struct {
	DiffFilePaths
	DiffStatistics
}

type ChunkPair struct {
	oldFileChunk Chunk
	newFileChunk Chunk
}

// Backend For Frontend
// ============Response body===============
type SplitDiff struct {
	DiffMetadata
	allChunkPairs []ChunkPair
	// Indices of changed chunk pair
	hunkPairIndices []int
}

type UnifiedDiff struct {
	DiffMetadata
	allChunks   []Chunk
	hunkIndices []int
}

// [l1, l2, l3, ...]
//

// 1		2		3		4		5		6
// line1, line2, line3, nothing, nothing, line4
// line1, line2, line3, line4,   line5,   line6

// oldChunk:
//	line1->1 line2->2 line3->3 nothing

// 0		 1		  2		   3		4		5
// line101, line102, line103, nothing, nothing, line104
// line201, line202, line203, line204, line205, line206

// lineMap [int]int // key: true line number, value: line index

// Unified View

// TODO: only see diff and hide unchanged
/*
Line 1: unchanged - -- +migrate Up
Line 2: deleted - CREATE TABLE amenity_type (
Line 2: added - CREATE TABLE amenity_type
Line 3: deleted - id VARCHAR(3) PRIMARY KEY,
Line 3: (
Line 4: deleted - title VARCHAR(100),
Line 4: added - id VARCHAR(3) PRIMARY KEY,



Diff:
	ChunkPair 1
	oldLines: Line2, Line3, Line4
	newLines: Line2, Line3, Line4, Line5


	ChunkPair 2
	oldLines: Line12, Line13, Line14
	newLines: Line13, Line14, Line15, Line16


func Display(diff Diff) {
	lineIndex int

	// for  chunkPair in range chunkPairs {
	// 	max := max(len(oldLines), len(newLines))
	// 	for i in range max {

	// 	}
	// }
}

*/
