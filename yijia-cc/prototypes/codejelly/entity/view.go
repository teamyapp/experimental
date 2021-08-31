package entity

// layer 5: feed data for split view and unified view
// Backend For Frontend
// ============Response body===============
type ChunkPair struct {
	FromFileChunk Chunk `json:"from_file_chunk"`
	ToFileChunk Chunk `json:"to_file_chunk"`
}

type SplitView struct {
	FileDiffMetadata FileDiffMetadata `json:"file_diff_metadata"`
	AllChunkPairs []ChunkPair `json:"all_chunk_pairs"`
	// Indices of changed chunk pair
	HunkPairIndices []int `json:"hunk_pair_indices"`
}

type UnifiedView struct {
	FileDiffMetadata FileDiffMetadata `json:"file_diff_metadata"`
	AllChunks   []Chunk `json:"all_chunks"`
	HunkIndices []int `json:"hunk_indices"`
}
