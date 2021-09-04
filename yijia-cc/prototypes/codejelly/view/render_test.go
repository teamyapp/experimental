package view

import (
	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/service"
	"testing"
)

func TestRenderUnifiedView(t *testing.T) {
	testCases := []struct {
		name string
		block string
		fileContent string
		expected entity.UnifiedView
		expectedHasErr bool
	} {
		{
			name: "empty file content",
			block:
`diff --git a/web/.env.development b/web/.env.development
new file mode 100644
index 0000000..edc67e8
--- /dev/null
+++ b/web/.env.development
@@ -0,0 +1 @@
+REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun
\ No newline at end of file`,
			fileContent: "",
			expected: entity.UnifiedView{
				FileDiffMetadata: entity.FileDiffMetadata{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeAdded,
						FromFilePath: "",
						ToFilePath: "web/.env.development",
						Similarity: 0,
					},
					FileDiffStatistics: entity.FileDiffStatistics{
						NumOfLinesChanged: 1,
						NumOfAddedLines: 1,
						NumOfDeletedLines: 0,
						FromFileTotalNumOfLines: 0,
						ToFileTotalNumOfLines: 1,
					},
				},
				AllChunks: []entity.Chunk{
					{
						NumberedLines: []entity.NumberedLine{
							{
								Status: entity.LineAdded,
								Content: "REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun",
								FromFileLineNumber: -1,
								ToFileLineNumber: 1,
							},
						},
						IsHunk: true,
					},
				},
				HunkIndices: []int{0},
			},
			expectedHasErr: false,
		},
		{
			name: "one hunk",
			block:
`diff --git a/calendar/gql/resolver/amenityinfo.go b/calendar/gql/resolver/amenityinfo.go
index 258d492..e51279e 100644
--- a/calendar/gql/resolver/amenityinfo.go
+++ b/calendar/gql/resolver/amenityinfo.go
@@ -3,7 +3,7 @@ type AmenityInfo struct {
 	amenityInfo entity.AmenityInfo
 }
 
-func (a AmenityInfo) Id() graphql.ID {
+func (a AmenityInfo) ID() graphql.ID {
 	return graphql.ID(a.amenityInfo.ID)
 }
 `,
 			fileContent:
`package resolver
type AmenityInfo struct {
	amenityInfo entity.AmenityInfo
}

func (a AmenityInfo) Id() graphql.ID {
	return graphql.ID(a.amenityInfo.ID)
}

func (a AmenityInfo) Name() *string {
	return &a.amenityInfo.Name
}`,
			expected: entity.UnifiedView{
				FileDiffMetadata: entity.FileDiffMetadata{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeModified,
						FromFilePath: "calendar/gql/resolver/amenityinfo.go",
						ToFilePath: "calendar/gql/resolver/amenityinfo.go",
						Similarity: 0,
					},
					FileDiffStatistics: entity.FileDiffStatistics{
						NumOfLinesChanged: 2,
						NumOfDeletedLines: 1,
						NumOfAddedLines: 1,
						FromFileTotalNumOfLines: 12,
						ToFileTotalNumOfLines: 12,
					},
				},
				AllChunks: []entity.Chunk{
					{
						NumberedLines: []entity.NumberedLine{
							{
								Status: entity.LineUnchanged,
								Content: "package resolver",
								FromFileLineNumber: 1,
								ToFileLineNumber: 1,
							},
							{
								Status: entity.LineUnchanged,
								Content: "type AmenityInfo struct {",
								FromFileLineNumber: 2,
								ToFileLineNumber: 2,
							},
						},
						IsHunk: false,
					},
					{
						NumberedLines: []entity.NumberedLine{
							{
								Status: entity.LineUnchanged,
								Content: "\tamenityInfo entity.AmenityInfo",
								FromFileLineNumber: 3,
								ToFileLineNumber: 3,
							},
							{
								Status: entity.LineUnchanged,
								Content: "}",
								FromFileLineNumber: 4,
								ToFileLineNumber: 4,
							},
							{
								Status: entity.LineUnchanged,
								Content: "",
								FromFileLineNumber: 5,
								ToFileLineNumber: 5,
							},
							{
								Status: entity.LineDeleted,
								Content: "func (a AmenityInfo) Id() graphql.ID {",
								FromFileLineNumber: 6,
								ToFileLineNumber: -1,
							},
							{
								Status: entity.LineAdded,
								Content: "func (a AmenityInfo) ID() graphql.ID {",
								FromFileLineNumber: -1,
								ToFileLineNumber: 6,
							},
							{
								Status: entity.LineUnchanged,
								Content: "\treturn graphql.ID(a.amenityInfo.ID)",
								FromFileLineNumber: 7,
								ToFileLineNumber: 7,
							},
							{
								Status: entity.LineUnchanged,
								Content: "}",
								FromFileLineNumber: 8,
								ToFileLineNumber: 8,
							},
							{
								Status: entity.LineUnchanged,
								Content: "",
								FromFileLineNumber: 9,
								ToFileLineNumber: 9,
							},
						},
						IsHunk: true,
					},
					{
						NumberedLines: []entity.NumberedLine{
							{
								Status: entity.LineUnchanged,
								Content: "func (a AmenityInfo) Name() *string {",
								FromFileLineNumber: 10,
								ToFileLineNumber: 10,
							},
							{
								Status: entity.LineUnchanged,
								Content: "\treturn &a.amenityInfo.Name",
								FromFileLineNumber: 11,
								ToFileLineNumber: 11,
							},
							{
								Status: entity.LineUnchanged,
								Content: "}",
								FromFileLineNumber: 12,
								ToFileLineNumber: 12,
							},
						},
						IsHunk: false,
					},
				},
				HunkIndices: []int{1},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T){
			t.Parallel()
			fileDiff, err := git.NewFileDiffFromBlock(testCase.block)
			assert.Nil(t, err)

			chunks := service.GetChunks(fileDiff.Hunks, testCase.fileContent)

			actual, err := RenderUnifiedView(entity.FullFileDiff{
				FileDiffHeader: fileDiff.FileDiffHeader,
				Chunks: chunks,
			})

			if testCase.expectedHasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestRenderSplitView(t *testing.T) {
	testCases := []struct {
		name string
		block string
		fileContent string
		expected entity.SplitView
		expectedHasErr bool
	} {
		{
			name: "empty file content",
			block:
			`diff --git a/web/.env.development b/web/.env.development
new file mode 100644
index 0000000..edc67e8
--- /dev/null
+++ b/web/.env.development
@@ -0,0 +1 @@
+REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun
\ No newline at end of file`,
			fileContent: "",
			expected: entity.SplitView{
				FileDiffMetadata: entity.FileDiffMetadata{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeAdded,
						FromFilePath: "",
						ToFilePath: "web/.env.development",
						Similarity: 0,
					},
					FileDiffStatistics: entity.FileDiffStatistics{
						NumOfLinesChanged: 1,
						NumOfAddedLines: 1,
						NumOfDeletedLines: 0,
						FromFileTotalNumOfLines: 0,
						ToFileTotalNumOfLines: 1,
					},
				},
				AllChunkPairs: []entity.ChunkPair{
					{
						FromFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{},
							IsHunk: true,
						},
						ToFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineAdded,
									Content: "REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun",
									FromFileLineNumber: -1,
									ToFileLineNumber: 1,
								},
							},
							IsHunk: true,
						},
					},
				},
				HunkPairIndices: []int{0},
			},
			expectedHasErr: false,
		},
		{
			name: "one hunk",
			block:
			`diff --git a/calendar/gql/resolver/amenityinfo.go b/calendar/gql/resolver/amenityinfo.go
index 258d492..e51279e 100644
--- a/calendar/gql/resolver/amenityinfo.go
+++ b/calendar/gql/resolver/amenityinfo.go
@@ -3,7 +3,7 @@ type AmenityInfo struct {
 	amenityInfo entity.AmenityInfo
 }
 
-func (a AmenityInfo) Id() graphql.ID {
+func (a AmenityInfo) ID() graphql.ID {
 	return graphql.ID(a.amenityInfo.ID)
 }
 `,
			fileContent:
			`package resolver
type AmenityInfo struct {
	amenityInfo entity.AmenityInfo
}

func (a AmenityInfo) Id() graphql.ID {
	return graphql.ID(a.amenityInfo.ID)
}

func (a AmenityInfo) Name() *string {
	return &a.amenityInfo.Name
}`,
			expected: entity.SplitView{
				FileDiffMetadata: entity.FileDiffMetadata{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeModified,
						FromFilePath: "calendar/gql/resolver/amenityinfo.go",
						ToFilePath: "calendar/gql/resolver/amenityinfo.go",
						Similarity: 0,
					},
					FileDiffStatistics: entity.FileDiffStatistics{
						NumOfLinesChanged: 2,
						NumOfDeletedLines: 1,
						NumOfAddedLines: 1,
						FromFileTotalNumOfLines: 12,
						ToFileTotalNumOfLines: 12,
					},
				},
				AllChunkPairs: []entity.ChunkPair{
					{
						FromFileChunk: entity.Chunk {
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "package resolver",
									FromFileLineNumber: 1,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "type AmenityInfo struct {",
									FromFileLineNumber: 2,
									ToFileLineNumber: -1,
								},
							},
							IsHunk: false,
						},
						ToFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "package resolver",
									FromFileLineNumber: -1,
									ToFileLineNumber: 1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "type AmenityInfo struct {",
									FromFileLineNumber: -1,
									ToFileLineNumber: 2,
								},
							},
							IsHunk: false,
						},
					},
					{
						FromFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "\tamenityInfo entity.AmenityInfo",
									FromFileLineNumber: 3,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: 4,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "",
									FromFileLineNumber: 5,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineDeleted,
									Content: "func (a AmenityInfo) Id() graphql.ID {",
									FromFileLineNumber: 6,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "\treturn graphql.ID(a.amenityInfo.ID)",
									FromFileLineNumber: 7,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: 8,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "",
									FromFileLineNumber: 9,
									ToFileLineNumber: -1,
								},
							},
							IsHunk: true,
						},
						ToFileChunk: entity.Chunk {
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "\tamenityInfo entity.AmenityInfo",
									FromFileLineNumber: -1,
									ToFileLineNumber: 3,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: -1,
									ToFileLineNumber: 4,
								},
								{
									Status: entity.LineUnchanged,
									Content: "",
									FromFileLineNumber: -1,
									ToFileLineNumber: 5,
								},
								{
									Status: entity.LineAdded,
									Content: "func (a AmenityInfo) ID() graphql.ID {",
									FromFileLineNumber: -1,
									ToFileLineNumber: 6,
								},
								{
									Status: entity.LineUnchanged,
									Content: "\treturn graphql.ID(a.amenityInfo.ID)",
									FromFileLineNumber: -1,
									ToFileLineNumber: 7,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: -1,
									ToFileLineNumber: 8,
								},
								{
									Status: entity.LineUnchanged,
									Content: "",
									FromFileLineNumber: -1,
									ToFileLineNumber: 9,
								},
							},
							IsHunk: true,
						},
					},
					{
						FromFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "func (a AmenityInfo) Name() *string {",
									FromFileLineNumber: 10,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "\treturn &a.amenityInfo.Name",
									FromFileLineNumber: 11,
									ToFileLineNumber: -1,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: 12,
									ToFileLineNumber: -1,
								},
							},
							IsHunk: false,
						},
						ToFileChunk: entity.Chunk{
							NumberedLines: []entity.NumberedLine{
								{
									Status: entity.LineUnchanged,
									Content: "func (a AmenityInfo) Name() *string {",
									FromFileLineNumber: -1,
									ToFileLineNumber: 10,
								},
								{
									Status: entity.LineUnchanged,
									Content: "\treturn &a.amenityInfo.Name",
									FromFileLineNumber: -1,
									ToFileLineNumber: 11,
								},
								{
									Status: entity.LineUnchanged,
									Content: "}",
									FromFileLineNumber: -1,
									ToFileLineNumber: 12,
								},
							},
							IsHunk: false,
						},
					},
				},
				HunkPairIndices: []int{1},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T){
			t.Parallel()
			fileDiff, err := git.NewFileDiffFromBlock(testCase.block)
			assert.Nil(t, err)

			chunks := service.GetChunks(fileDiff.Hunks, testCase.fileContent)

			actual, err := RenderSplitView(entity.FullFileDiff{
				FileDiffHeader: fileDiff.FileDiffHeader,
				Chunks: chunks,
			})

			if testCase.expectedHasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.expected, actual)
		})
	}
}

