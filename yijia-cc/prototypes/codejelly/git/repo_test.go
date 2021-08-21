package git_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git/gittest"
	"testing"
)

var outputMap = map[string]string{
"feature1,master": "",
"feature2,master":
`
A       dashboard/model/event.go
D       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
"feature3,master":
`

M       discussion/.gitignore +
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
"master,feature2":
`
D       dashboard/model/event.go
A       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/UserModel.java discussion/src/main/java/info/User.java`,
"feature4,master":
`
A       dashboard/model/event.go
B       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
}

func TestGetFileDiffsBetweenBranches(t *testing.T) {
	testCases := []struct {
		name        string
		fromBranch  string
		toBranch    string
		expected    []entity.FileDiffHeader
		expectedHasErr bool
	}{
		{
			name:        "no change",
			fromBranch:  "feature1",
			toBranch:    "master",
			expected:    nil,
			expectedHasErr: true,
		},
		{
			name:       "valid command output",
			fromBranch: "feature2",
			toBranch:   "master",
			expected: []entity.FileDiffHeader{
				{
					Status:       entity.ChangeAdded,
					FromFilePath: "dashboard/model/event.go",
					ToFilePath:   "dashboard/model/event.go",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeDeleted,
					FromFilePath: "dashboard/seeds/seeder.go",
					ToFilePath:   "dashboard/seeds/seeder.go",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeModified,
					FromFilePath: "discussion/.gitignore",
					ToFilePath:   "discussion/.gitignore",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeRenamed,
					FromFilePath: "discussion/src/main/java/info/User.java",
					ToFilePath:   "discussion/src/main/java/info/UserModel.java",
					Similarity:   58,
				},
			},
			expectedHasErr: false,
		},
		{
			name:       "contains empty line",
			fromBranch: "feature3",
			toBranch:   "master",
			expected: []entity.FileDiffHeader{
				{
					Status:       entity.ChangeModified,
					FromFilePath: "discussion/.gitignore",
					ToFilePath:   "discussion/.gitignore",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeRenamed,
					FromFilePath: "discussion/src/main/java/info/User.java",
					ToFilePath:   "discussion/src/main/java/info/UserModel.java",
					Similarity:   58,
				},
			},
			expectedHasErr: false,
		},
		{
			name:        "command output contains invalid change status",
			fromBranch:  "feature4",
			toBranch:    "master",
			expected:    nil,
			expectedHasErr: true,
		},
		{
			name:       "switch fromBranch and toBranch",
			fromBranch: "master",
			toBranch:   "feature2",
			expected: []entity.FileDiffHeader{
				{
					Status:       entity.ChangeDeleted,
					FromFilePath: "dashboard/model/event.go",
					ToFilePath:   "dashboard/model/event.go",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeAdded,
					FromFilePath: "dashboard/seeds/seeder.go",
					ToFilePath:   "dashboard/seeds/seeder.go",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeModified,
					FromFilePath: "discussion/.gitignore",
					ToFilePath:   "discussion/.gitignore",
					Similarity:   0,
				},
				{
					Status:       entity.ChangeRenamed,
					FromFilePath: "discussion/src/main/java/info/UserModel.java",
					ToFilePath:   "discussion/src/main/java/info/User.java",
					Similarity:   58,
				},
			},
			expectedHasErr: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			stubCommandExecutor := gittest.NewStubCommandExecutor(outputMap)
			repo := git.NewGitCustomExecutor(stubCommandExecutor, "/repo/")
			actual, err := repo.GetFileDiffHeadersBetweenBranches(testCase.fromBranch, testCase.toBranch)
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
