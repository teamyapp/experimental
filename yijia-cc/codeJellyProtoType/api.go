
type PullRequest struct {
	id int
    title string
	description string
	author User
	reviewers []Reviewer

	repo string	
	sourceBranch string
	targetBranch string
	parentId *int // -1 root PR/ base PR
	openedAt time.Time
	availableActions []PullRequestAction
	
	pullRequestStatus PullRequestStatus
	comments []Comment
	commits []Commit
	aggregatedFileChanges 
}

// Policy

type PullRequestStatus int 

const (
	opened PullRequestStatus = iota
	//unresolved
	approved
	merged
	closed
)

type PullRequestAction int 

const (
	approve PullRequestAction = iota
	close
	merge
	squashAndMerge
	rebaseAndMerge
)

type Thread struct {
	id int
	comments []Comment
	// mark unresolved when new comments added
	isResolved bool 
	selection Selection
}

type Comment struct {
	id int
	reviewerId int
	// TODO: 
	// 1) allow emojis 
	// 2) allow tagging others 
	// 3) quote code 
	// 4) offer code suggestions
	content string 
	createdAt time.Time
}

type Selection struct {
	startLine int
	startColumn int
	endLine int
	endColumn int
}

// type Diff struct {
// 	chunkPairs []ChunkPairs
// }

type ChunkPairs struct {
	oldLines []Line
	newLines []Line
}

type Line struct {
	number int
	status LineStatus
	content string
}

type LineStatus int

const (
	unchanged LineStatus = iota 
	deleted 
	added
	nothing
)

type Reviewer struct {
	user User
	reviewStatus ReviewStatus
}

type ReviewStatus int

const (
	unviewed ReviewerStatus = iota
	unresolvedComment
	approved
)

type User struct {
	id int
	name string
}

type Commit struct {
	id int
	message string
	createdAt time.Time
	changedFiles []ChangedFile
	author User
}

type ChangedFile struct {
	id int
	filePath string
	commitId int
	changeStatus ChangeStatus
	numOfChangedLines int
	comments []Comment
}

type ChangeStatus int

const (
	added ChangeStatus = iota
	modified
	deleted
)


type Diff struct {
	hunks []Hunk
}

type Hunk struct {
	fromFilePath string
	toFilePath string
	similarity int
	renameFromFilePath *string
	renameToFilePath *string
	//index string

	fromFileStartLine int
	fromFileNumOfLines int
	toFileStartLine int
	toFileNumOfLines int
	
	lines []Line // - line 16-37, + line 16-33, unchanged
}

type CodeChunkPair {
	fromFileCodeChunk CodeChunk
	toFileCodeChunk CodeChunk
}

type CodeChunk struct {
	startLine int
	endLine int
	lines []Line
	lineMap [int]int // key: true line number, value: line index
	fileSource string // fromFile / toFile

	isHunk bool
	
}

type File struct {
	CodeChunks []CodeChunk
}



/*
layer 1: git diff
layer 2: unorganized hunks
layer 3: group unchaned chunks and hunks into file change pair
layer 4: feed data for split view and unified view
layer 5: render UI at frontend
*/



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

