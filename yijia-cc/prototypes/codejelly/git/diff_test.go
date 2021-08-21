package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func TestNewFileDiffFromLine(t *testing.T) {
	testCases := []struct {
		name string
		input string
		expected entity.FileDiffHeader
		expectedHasErr bool
	} {
		{
			name: "empty line",
			input: "",
			expected: entity.FileDiffHeader{},
			expectedHasErr: false,
		},
		{
			name: "contains status only",
			input: "M",
			expected: entity.FileDiffHeader{},
			expectedHasErr: false,
		},
		{
			name: "unsupported status",
			input: "B dashboard/model/db.go",
			expected: entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name: "similarity not number",
			input: "RTT    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name: "added 1 line",
			input: "A       dashboard/model/db.go",
			expected: entity.FileDiffHeader{
				Status: entity.ChangeAdded,
				FromFilePath: "dashboard/model/db.go",
				ToFilePath: "dashboard/model/db.go",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "deleted 1 line",
			input: "D       calendar/repo/schedule.go",
			expected: entity.FileDiffHeader{
				Status: entity.ChangeDeleted,
				FromFilePath: "calendar/repo/schedule.go",
				ToFilePath: "calendar/repo/schedule.go",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "modified 1 line",
			input: "M       discussion/pom.xml",
			expected: entity.FileDiffHeader{
				Status: entity.ChangeModified,
				FromFilePath: "discussion/pom.xml",
				ToFilePath: "discussion/pom.xml",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "renamed a file",
			input: "R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiffHeader{
				Status: entity.ChangeRenamed,
				FromFilePath: "discussion/src/main/java/info/User.java",
				ToFilePath: "discussion/src/main/java/info/UserModel.java",
				Similarity: 58,
			},
			expectedHasErr: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := newFileDiffHeaderFromLine(testCase.input)
			if testCase.expectedHasErr {
				assert.NotNil(t, err)
				return
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
