package entity

//layer 3: core data model:
//group unchanged chunks and hunks into file change pair


type FullFileDiff struct {
	FileDiffHeader
	Chunks []Chunk
}

