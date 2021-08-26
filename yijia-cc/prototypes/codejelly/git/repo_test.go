package git_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git/gittest"
)

var fileDiffOutputMap = map[string]string{
	"feature1,master": "",
	"feature2,master": `diff --git a/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java b/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
index 3d739dc..40b8555 100644
--- a/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
+++ b/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
@@ -9,7 +9,8 @@ import info.grouplive.discussion.exceptions.PostNotFoundException;
 import info.grouplive.discussion.mapper.CommentMapper;
-import info.grouplive.discussion.model.User;
-//import info.grouplive.discussion.model.User;
+import info.grouplive.discussion.model.UserModel;
 import lombok.AllArgsConstructor;
@@ -52,7 +45,7 @@ public class CommentService {
     }
 
     public List<CommentsDto> getAllCommentsForUser(String userName) {
-        User user = userRepository.findByUsername(userName)
+        UserModel user = userRepository.findByUsername(userName)
                             .orElseThrow(() -> new UsernameNotFoundException(userName));
         return commentRepository.findAllByUser(user)
                 .stream()
diff --git a/web/.env.development b/web/.env.development
new file mode 100644
index 0000000..edc67e8
--- /dev/null
+++ b/web/.env.development
@@ -0,0 +1 @@
+REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun
\ No newline at end of file`,
	"feature3,master":`diff --git a/src/main/resources/static/img/addPics.png b/src/main/resources/static/img/addPics.png
added file mode 100755
index 3ba6527..0000000
Binary files /dev/null and b/src/main/resources/static/img/addPics.png differ`,
	"feature4,master":`diff --git a/src/main/resources/static/img/default_avatar.jpg b/src/main/resources/static/img/default_avatar.jpeg
similarity index 100%
rename from src/main/resources/static/img/default_avatar.jpg
rename to src/main/resources/static/img/default_avatar.jpeg`,
	"feature5,master":`diff --git a/src/main/resources/static/img/addPics.png b/src/main/resources/static/img/addPics.png
deleted file mode 100755
index 3ba6527..0000000
Binary files a/src/main/resources/static/img/addPics.png and /dev/null differ`,
	"feature6,master":`diff --git a/src/main/resources/static/img/avatar.png b/src/main/resources/static/img/avatar.png
old mode 100755
new mode 100644
index a6a5318..495eafd
Binary files a/src/main/resources/static/img/avatar.png and b/src/main/resources/static/img/avatar.png differ`,
}

var fileDiffHeaderOutputMap = map[string]string{


	"feature1,master": "",
	"feature2,master": `
A       dashboard/model/event.go
D       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
	"feature3,master": `

M       discussion/.gitignore +
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
	"master,feature2": `
D       dashboard/model/event.go
A       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/UserModel.java discussion/src/main/java/info/User.java`,
	"feature4,master": `
A       dashboard/model/event.go
B       dashboard/seeds/seeder.go
M       discussion/.gitignore
R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java`,
}

func TestGetFileDiffHeadersBetweenBranches(t *testing.T) {
	testCases := []struct {
		name           string
		fromBranch     string
		toBranch       string
		expected       []entity.FileDiffHeader
		expectedHasErr bool
	}{
		{
			name:           "no change",
			fromBranch:     "feature1",
			toBranch:       "master",
			expected:       nil,
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
			name:           "command output contains invalid change status",
			fromBranch:     "feature4",
			toBranch:       "master",
			expected:       nil,
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

			stubCommandExecutor := gittest.NewStubCommandExecutor(fileDiffHeaderOutputMap)
			repo := git.NewRepositoryDeps(stubCommandExecutor, "/repo/")
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

func TestParseFileDiffsFromOutput(t *testing.T) {
	testCases := []struct {
		name           string
		fromBranch     string
		toBranch       string
		expected       []entity.FileDiff
		expectedHasErr bool
	}{
		{
			name:           "empty block",
			fromBranch:     "feature1",
			toBranch:       "master",
			expected:       nil,
			expectedHasErr: true,
		},
		{
			name:       "file modified with multiple hunks",
			fromBranch: "feature2",
			toBranch:   "master",
			expected: []entity.FileDiff{
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status:       entity.ChangeModified,
						FromFilePath: "discussion/src/main/java/info/grouplive/discussion/service/CommentService.java",
						ToFilePath:   "discussion/src/main/java/info/grouplive/discussion/service/CommentService.java",
						Similarity:   0,
					},
					Hunks: []entity.Hunk{
						{
							HunkHeader: entity.HunkHeader{
								FromFileStartLine:  9,
								FromFileNumOfLines: 7,
								ToFileStartLine:    9,
								ToFileNumOfLines:   8,
							},

							Lines: []entity.Line{
								{
									Status: entity.LineHunkHeader,
									Content: "@@ -9,7 +9,8 @@ import info.grouplive.discussion.exceptions.PostNotFoundException;",
								},
								{
									Status:  entity.LineUnchanged,
									Content: " import info.grouplive.discussion.mapper.CommentMapper;",
								},
								{
									Status:  entity.LineDeleted,
									Content: "-import info.grouplive.discussion.model.User;",
								},
								{
									Status:  entity.LineDeleted,
									Content: "-//import info.grouplive.discussion.model.User;",
								},
								{
									Status:  entity.LineAdded,
									Content: "+import info.grouplive.discussion.model.UserModel;",
								},
								{
									Status:  entity.LineUnchanged,
									Content: " import lombok.AllArgsConstructor;",
								},
							},
						},
						{
							HunkHeader: entity.HunkHeader{
								FromFileStartLine:  52,
								FromFileNumOfLines: 7,
								ToFileStartLine:    45,
								ToFileNumOfLines:   7,
							},

							Lines: []entity.Line{
								{
									Status: entity.LineHunkHeader,
									Content: "@@ -52,7 +45,7 @@ public class CommentService {",
								},
								{
									Status:  entity.LineUnchanged,
									Content: "     }",
								},
								{
									Status:  entity.LineUnchanged,
									Content: " ",
								},
								{
									Status:  entity.LineUnchanged,
									Content: "     public List<CommentsDto> getAllCommentsForUser(String userName) {",
								},
								{
									Status:  entity.LineDeleted,
									Content: "-        User user = userRepository.findByUsername(userName)",
								},
								{
									Status:  entity.LineAdded,
									Content: "+        UserModel user = userRepository.findByUsername(userName)",
								},
								{
									Status:  entity.LineUnchanged,
									Content: "                             .orElseThrow(() -> new UsernameNotFoundException(userName));",
								},
								{
									Status:  entity.LineUnchanged,
									Content: "         return commentRepository.findAllByUser(user)",
								},
								{
									Status:  entity.LineUnchanged,
									Content: "                 .stream()",
								},
							},
						},
					},
				},
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status:       entity.ChangeAdded,
						FromFilePath: "/dev/null",
						ToFilePath:   "web/.env.development",
						Similarity:   0,
					},
					Hunks: []entity.Hunk{
						{
							HunkHeader: entity.HunkHeader{
								FromFileStartLine:  0,
								FromFileNumOfLines: 0,
								ToFileStartLine:    1,
								ToFileNumOfLines:   1,
							},

							Lines: []entity.Line{
								{
									Status: entity.LineHunkHeader,
									Content: "@@ -0,0 +1 @@",
								},
								{
									Status:  entity.LineAdded,
									Content: "+REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "added binary file",
			fromBranch: "feature3",
			toBranch: "master",
			expected: []entity.FileDiff{
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeAdded,
						FromFilePath: "/dev/null",
						ToFilePath: "src/main/resources/static/img/addPics.png",
						Similarity: 0,
					},
					Hunks: []entity.Hunk{},
				},
			},
		},
		{
			name: "renamed binary file",
			fromBranch: "feature4",
			toBranch: "master",
			expected: []entity.FileDiff{
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeRenamed,
						FromFilePath: "src/main/resources/static/img/default_avatar.jpg",
						ToFilePath: "src/main/resources/static/img/default_avatar.jpeg",
						Similarity: 100,
					},
					Hunks: []entity.Hunk{},
				},
			},
		},
		{
			name: "deleted binary file",
			fromBranch: "feature5",
			toBranch: "master",
			expected: []entity.FileDiff{
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeDeleted,
						FromFilePath: "src/main/resources/static/img/addPics.png",
						ToFilePath: "/dev/null",
						Similarity: 0,
					},
					Hunks: []entity.Hunk{},
				},
			},
		},
		{
			name: "modified binary file",
			fromBranch: "feature6",
			toBranch: "master",
			expected: []entity.FileDiff{
				{
					FileDiffHeader: entity.FileDiffHeader{
						Status: entity.ChangeModified,
						FromFilePath: "src/main/resources/static/img/avatar.png",
						ToFilePath: "src/main/resources/static/img/avatar.png",
						Similarity: 0,
					},
					Hunks: []entity.Hunk{},
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			stubCommandExecutor := gittest.NewStubCommandExecutor(fileDiffOutputMap)
			repo := git.NewRepositoryDeps(stubCommandExecutor, "/repo/")
			actual, err := repo.GetFileDiffsBetweenBranches(testCase.fromBranch, testCase.toBranch)
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

