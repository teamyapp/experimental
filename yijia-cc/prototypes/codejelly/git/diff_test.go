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
		expectedErr bool
	} {
		{
			name: "empty line",
			input: "",
			expected: entity.FileDiffHeader{},
			expectedErr: false,
		},
		{
			name: "contains status only",
			input: "M",
			expected: entity.FileDiffHeader{},
			expectedErr: false,
		},
		{
			name: "unsupported status",
			input: "B dashboard/model/db.go",
			expected: entity.FileDiffHeader{},
			expectedErr: true,
		},
		{
			name: "similarity not number",
			input: "RTT    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiffHeader{},
			expectedErr: true,
		},
		{
			name: "added 1 line",
			input: "A       dashboard/model/db.go",
			expected: entity.FileDiffHeader{
				Id: 0,
				Status: entity.ChangeAdded,
				FromFilePath: "dashboard/model/db.go",
				ToFilePath: "dashboard/model/db.go",
				Similarity: 0,
			},
			expectedErr: false,
		},
		{
			name: "deleted 1 line",
			input: "D       calendar/repo/schedule.go",
			expected: entity.FileDiffHeader{
				Id: 0,
				Status: entity.ChangeDeleted,
				FromFilePath: "calendar/repo/schedule.go",
				ToFilePath: "calendar/repo/schedule.go",
				Similarity: 0,
			},
			expectedErr: false,
		},
		{
			name: "modified 1 line",
			input: "M       discussion/pom.xml",
			expected: entity.FileDiffHeader{
				Id: 0,
				Status: entity.ChangeModified,
				FromFilePath: "discussion/pom.xml",
				ToFilePath: "discussion/pom.xml",
				Similarity: 0,
			},
			expectedErr: false,
		},
		{
			name: "renamed a file",
			input: "R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiffHeader{
				Id: 0,
				Status: entity.ChangeRenamed,
				FromFilePath: "discussion/src/main/java/info/User.java",
				ToFilePath: "discussion/src/main/java/info/UserModel.java",
				Similarity: 58,
			},
			expectedErr: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := newFileDiffHeaderFromLine(testCase.input, 0)
			if testCase.expectedErr && err != nil{
				assert.NotNil(t, err)
				return
			}

			if testCase.expectedErr || err != nil {
				t.Fail()
			}

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
