package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"testing"
)

func TestGetChunks(t *testing.T) {
	testCases := []struct {
		name string
		hunks []entity.Hunk
		fromFileLines []string
		expected []entity.Chunk
	} {
		{
			name: "multiple hunks",
			hunks: []entity.Hunk{
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine: 2,
						FromFileNumOfLines: 4,
						ToFileStartLine: 2,
						ToFileNumOfLines: 3,
					},

					Lines: []entity.Line{
						{
							Status: entity.LineUnchanged,
							Content: "import info.grouplive.discussion.mapper.CommentMapper;",
						},
						{
							Status: entity.LineDeleted,
							Content: "import info.grouplive.discussion.model.User;",
						},
						{
							Status: entity.LineDeleted,
							Content: "//import info.grouplive.discussion.model.User;",
						},
						{
							Status: entity.LineAdded,
							Content: "import info.grouplive.discussion.model.UserModel;",
						},
						{
							Status: entity.LineUnchanged,
							Content: "import lombok.AllArgsConstructor;",
						},
					},
				},
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine: 8,
						FromFileNumOfLines: 4,
						ToFileStartLine: 7,
						ToFileNumOfLines: 4,
					},

					Lines: []entity.Line{
						{
							Status: entity.LineUnchanged,
							Content: "    }",
						},
						{
							Status: entity.LineUnchanged,
							Content: "",
						},
						{
							Status: entity.LineUnchanged,
							Content: "    public List<CommentsDto> getAllCommentsForUser(String userName) {",
						},
						{
							Status: entity.LineDeleted,
							Content: "        User user = userRepository.findByUsername(userName)",
						},
						{
							Status: entity.LineAdded,
							Content: "        UserModel user = userRepository.findByUsername(userName)",
						},
					},
				},
			},
			fromFileLines: []string{
				"import info.grouplive.discussion.exceptions.PostNotFoundException;",
				"import info.grouplive.discussion.mapper.CommentMapper;",
				"import info.grouplive.discussion.model.User;",
				"//import info.grouplive.discussion.model.User;",
				"import lombok.AllArgsConstructor;",
				"",
				"public class CommentService {",
				"    }",
				"",
				"    public List<CommentsDto> getAllCommentsForUser(String userName) {",
				"        User user = userRepository.findByUsername(userName)",
			},
			expected: []entity.Chunk {
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status:  entity.LineUnchanged,
							Content: "import info.grouplive.discussion.exceptions.PostNotFoundException;",
							FromFileLineNumber: 1,
							ToFileLineNumber: 1,
						},
					},
					IsHunk: false,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "import info.grouplive.discussion.mapper.CommentMapper;",
							FromFileLineNumber: 2,
							ToFileLineNumber: 2,
						},
						{
							Status: entity.LineDeleted,
							Content: "import info.grouplive.discussion.model.User;",
							FromFileLineNumber: 3,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineDeleted,
							Content: "//import info.grouplive.discussion.model.User;",
							FromFileLineNumber: 4,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineAdded,
							Content: "import info.grouplive.discussion.model.UserModel;",
							FromFileLineNumber: entity.NoLineNumber,
							ToFileLineNumber: 3,
						},
						{
							Status: entity.LineUnchanged,
							Content: "import lombok.AllArgsConstructor;",
							FromFileLineNumber: 5,
							ToFileLineNumber: 4,
						},
					},
					IsHunk: true,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 6,
							ToFileLineNumber: 5,
						},
						{
							Status: entity.LineUnchanged,
							Content: "public class CommentService {",
							FromFileLineNumber: 7,
							ToFileLineNumber: 6,
						},
					},
					IsHunk: false,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "    }",
							FromFileLineNumber: 8,
							ToFileLineNumber: 7,
						},
						{
							Status: entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 9,
							ToFileLineNumber: 8,
						},
						{
							Status: entity.LineUnchanged,
							Content: "    public List<CommentsDto> getAllCommentsForUser(String userName) {",
							FromFileLineNumber: 10,
							ToFileLineNumber: 9,
						},
						{
							Status: entity.LineDeleted,
							Content: "        User user = userRepository.findByUsername(userName)",
							FromFileLineNumber: 11,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineAdded,
							Content: "        UserModel user = userRepository.findByUsername(userName)",
							FromFileLineNumber: entity.NoLineNumber,
							ToFileLineNumber: 10,
						},
					},
					IsHunk: true,
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := getChunks(testCase.hunks, testCase.fromFileLines)

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
