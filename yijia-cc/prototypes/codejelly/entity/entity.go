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

	fullFileDiffs []FullFileDiff
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
complexity:
- tech stack
	- system design
	- git/github
	- go
	- db
	- caching
	- concurrency
	- distributed storage
- learn from the project
	- git command
- problem solving difficulty
- product impact (long term/ short term)
- extensibility

layer 1: git diff
layer 2: unorganized hunks
layer 3: group unchanged chunks and hunks into file change pair
- contains all data that is needed in layer 4 and layer 5 now and in the future

layer 4: compute statistics
layer 5: feed data for split view and unified view
layer 6: Web APIs, gRPC, GraphQL
layer 7: render UI at frontend with React

//TODO: show code change in file side by side


*/

// Web API
// React App: TypeScript + SCSS, Storybook
// Individual component


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


//type FilePaths struct {
//	oldFilePath string
//	newFilePath string
//	isRenamed   bool
//}


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
