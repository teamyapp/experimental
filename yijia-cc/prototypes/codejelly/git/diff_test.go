package git

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func TestNewFileDiffFromLine(t *testing.T) {
	testCases := []struct {
		name string
		input string
		expected entity.FileDiff
		expectedErr error
	} {
		{
			name: "Empty line",
			input: "",
			expected: entity.FileDiff{},
			expectedErr: nil,
		},
		{
			name: "Contains only one part",
			input: "AAAAAAAAAAA",
			expected: entity.FileDiff{},
			expectedErr: nil,
		},
		{
			name: "Invalid Status",
			input: "B dashboard/model/db.go",
			expected: entity.FileDiff{},
			expectedErr: errors.New("invalid status"),
		},
		{
			name: "Renamed line invalid similarity",
			input: "RTT    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiff{},
			expectedErr: errors.New("invalid similarity"),
		},
		{
			name: "Added line",
			input: "A       dashboard/model/db.go",
			expected: entity.FileDiff{
				Status: entity.ChangeAdded,
				FromFilePath: "dashboard/model/db.go",
				ToFilePath: "dashboard/model/db.go",
				Similarity: 0,
			},
			expectedErr: nil,
		},
		{
			name: "Deleted line",
			input: "D       calendar/repo/schedule.go",
			expected: entity.FileDiff{
				Status: entity.ChangeDeleted,
				FromFilePath: "calendar/repo/schedule.go",
				ToFilePath: "calendar/repo/schedule.go",
				Similarity: 0,
			},
			expectedErr: nil,
		},
		{
			name: "Modified line",
			input: "M       discussion/pom.xml",
			expected: entity.FileDiff{
				Status: entity.ChangeModified,
				FromFilePath: "discussion/pom.xml",
				ToFilePath: "discussion/pom.xml",
				Similarity: 0,
			},
			expectedErr: nil,
		},
		{
			name: "Renamed line",
			input: "R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiff{
				Status: entity.ChangeRenamed,
				FromFilePath: "discussion/src/main/java/info/User.java",
				ToFilePath: "discussion/src/main/java/info/UserModel.java",
				Similarity: 58,
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := newFileDiffFromLine(testCase.input)
			fmt.Println(actual)
			if testCase.expectedErr != nil && err != nil{
				assert.NotNil(t, err)
				return
			}

			if testCase.expectedErr != nil || err != nil {
				t.Fail()
			}

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
