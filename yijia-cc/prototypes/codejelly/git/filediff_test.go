package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

func TestNewFileDiffHeaderFromLine(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expected       entity.FileDiffHeader
		expectedHasErr bool
	}{
		{
			name:           "empty line",
			input:          "",
			expected:       entity.FileDiffHeader{},
			expectedHasErr: false,
		},
		{
			name:           "contains status only",
			input:          "M",
			expected:       entity.FileDiffHeader{},
			expectedHasErr: false,
		},
		{
			name:           "unsupported status",
			input:          "B dashboard/model/db.go",
			expected:       entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name:           "similarity not number",
			input:          "RTT    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected:       entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name:  "added 1 line",
			input: "A       dashboard/model/db.go",
			expected: entity.FileDiffHeader{
				Status:       entity.ChangeAdded,
				FromFilePath: "dashboard/model/db.go",
				ToFilePath:   "dashboard/model/db.go",
				Similarity:   0,
			},
			expectedHasErr: false,
		},
		{
			name:  "deleted 1 line",
			input: "D       calendar/repo/schedule.go",
			expected: entity.FileDiffHeader{
				Status:       entity.ChangeDeleted,
				FromFilePath: "calendar/repo/schedule.go",
				ToFilePath:   "calendar/repo/schedule.go",
				Similarity:   0,
			},
			expectedHasErr: false,
		},
		{
			name:  "modified 1 line",
			input: "M       discussion/pom.xml",
			expected: entity.FileDiffHeader{
				Status:       entity.ChangeModified,
				FromFilePath: "discussion/pom.xml",
				ToFilePath:   "discussion/pom.xml",
				Similarity:   0,
			},
			expectedHasErr: false,
		},
		{
			name:  "renamed a file",
			input: "R058    discussion/src/main/java/info/User.java discussion/src/main/java/info/UserModel.java",
			expected: entity.FileDiffHeader{
				Status:       entity.ChangeRenamed,
				FromFilePath: "discussion/src/main/java/info/User.java",
				ToFilePath:   "discussion/src/main/java/info/UserModel.java",
				Similarity:   58,
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

func TestNewHunkFromBlock(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expected       []entity.Hunk
		expectedHasErr bool
	}{
		{
			name:           "empty block",
			input:          "",
			expected:       nil,
			expectedHasErr: false,
		},
		{
			name: "hunk with only one line added",
			input: `
a/web/.env.development b/web/.env.development
new file mode 100644
index 0000000..edc67e8
--- /dev/null
+++ b/web/.env.development
@@ -0,0 +1 @@
+REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun
\ No newline at end of file`,
			expected: []entity.Hunk{
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine: 0,
						FromFileNumOfLines: 0,
						ToFileStartLine: 1,
						ToFileNumOfLines: 1,
						HeaderLine: "",
					},
					Lines: []entity.Line{
						{
							Status: entity.LineAdded,
							Content: "REACT_APP_AUTH_API_BASE_URL=http://auth.api.staging.allgame.fun",
						},
					},
				},
			},
			expectedHasErr: false,
		},
		{
			name: "hunk with only one line deleted",
			input: `
a/src/main/java/com/yijia/dianping/model/TestModel.java b/src/main/java/com/yijia/dianping/model/TestModel.java
deleted file mode 100644
index 345e6ae..0000000
--- a/src/main/java/com/yijia/dianping/model/TestModel.java
+++ /dev/null
@@ -1 +0,0 @@
-Test`,
			expected: []entity.Hunk{
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine: 1,
						FromFileNumOfLines: 1,
						ToFileStartLine: 0,
						ToFileNumOfLines: 0,
						HeaderLine: "",
					},
					Lines: []entity.Line{
						{
							Status: entity.LineDeleted,
							Content: "Test",
						},
					},
				},
			},
			expectedHasErr: false,
		},
		{
			name: "invalid stats string",
			input: `
a/src/main/java/com/yijia/dianping/model/TestModel.java b/src/main/java/com/yijia/dianping/model/TestModel.java
deleted file mode 100644
index 345e6ae..0000000
--- a/src/main/java/com/yijia/dianping/model/TestModel.java
+++ /dev/null
@@ xx @@
-Test`,
			expected: nil,
			expectedHasErr: true,
		},
		{
			name: "multiple hunks",
			input: `
 a/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java b/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
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
                 .stream()`,
			expected: []entity.Hunk{
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine:  9,
						FromFileNumOfLines: 7,
						ToFileStartLine:    9,
						ToFileNumOfLines:   8,
						HeaderLine: "import info.grouplive.discussion.exceptions.PostNotFoundException;",
					},

					Lines: []entity.Line{
						{
							Status:  entity.LineUnchanged,
							Content: "import info.grouplive.discussion.mapper.CommentMapper;",
						},
						{
							Status:  entity.LineDeleted,
							Content: "import info.grouplive.discussion.model.User;",
						},
						{
							Status:  entity.LineDeleted,
							Content: "//import info.grouplive.discussion.model.User;",
						},
						{
							Status:  entity.LineAdded,
							Content: "import info.grouplive.discussion.model.UserModel;",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "import lombok.AllArgsConstructor;",
						},
					},
				},
				{
					HunkHeader: entity.HunkHeader{
						FromFileStartLine:  52,
						FromFileNumOfLines: 7,
						ToFileStartLine:    45,
						ToFileNumOfLines:   7,
						HeaderLine: "public class CommentService {",
					},

					Lines: []entity.Line{
						{
							Status:  entity.LineUnchanged,
							Content: "    }",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "    public List<CommentsDto> getAllCommentsForUser(String userName) {",
						},
						{
							Status:  entity.LineDeleted,
							Content: "        User user = userRepository.findByUsername(userName)",
						},
						{
							Status:  entity.LineAdded,
							Content: "        UserModel user = userRepository.findByUsername(userName)",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "                            .orElseThrow(() -> new UsernameNotFoundException(userName));",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "        return commentRepository.findAllByUser(user)",
						},
						{
							Status:  entity.LineUnchanged,
							Content: "                .stream()",
						},
					},
				},
			},
			expectedHasErr: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := newHunkFromBlock(testCase.input)
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

func TestNewFileDiffHeaderFromBlock(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expected       entity.FileDiffHeader
		expectedHasErr bool
	}{
		{
			name:           "empty block",
			input:          "",
			expected:       entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name: "file modified with multiple hunks",
			input: `
 a/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java b/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
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
                 .stream()`,
			expected: entity.FileDiffHeader{
				Status:       entity.ChangeModified,
				FromFilePath: "discussion/src/main/java/info/grouplive/discussion/service/CommentService.java",
				ToFilePath:   "discussion/src/main/java/info/grouplive/discussion/service/CommentService.java",
				Similarity:   0,
			},
			expectedHasErr: false,
		},
		{
			name: "file added",
			input: `
a/discussion/src/main/resources/.env_local b/discussion/src/main/resources/.env_local
new file mode 100644
index 0000000..dbc2c89
--- /dev/null
+++ b/discussion/src/main/resources/.env_local`,
			expected: entity.FileDiffHeader{
				Status: entity.ChangeAdded,
				FromFilePath: "/dev/null",
				ToFilePath: "discussion/src/main/resources/.env_local",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "file deleted",
			input: `
a/dashboard/controller/controller.go b/dashboard/controller/controller.go
deleted file mode 100644
index 015c776..0000000
--- a/dashboard/controller/controller.go
+++ /dev/null`,
			expected: entity.FileDiffHeader{
				Status: entity.ChangeDeleted,
				FromFilePath: "dashboard/controller/controller.go",
				ToFilePath: "/dev/null",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "file renamed",
			input: `
a/discussion/src/main/java/info/grouplive/discussion/model/User.java b/discussion/src/main/java/info/grouplive/discussion/model/UserModel.java
similarity index 58%
rename from discussion/src/main/java/info/grouplive/discussion/model/User.java
rename to discussion/src/main/java/info/grouplive/discussion/model/UserModel.java
index fc5ce3e..9b2a4a5 100644
--- a/discussion/src/main/java/info/grouplive/discussion/model/User.java
+++ b/discussion/src/main/java/info/grouplive/discussion/model/UserModel.java`,
			expected: entity.FileDiffHeader{
				Status: entity.ChangeRenamed,
				FromFilePath: "discussion/src/main/java/info/grouplive/discussion/model/User.java",
				ToFilePath: "discussion/src/main/java/info/grouplive/discussion/model/UserModel.java",
				Similarity: 58,
			},
			expectedHasErr: false,
		},
		{
			name: "binary file modified",
			input: `
a/proto/java/build/libs/proto-1.0-SNAPSHOT.jar b/proto/java/build/libs/proto-1.0-SNAPSHOT.jar
index 9471a5c..e09833d 100644
Binary files a/proto/java/build/libs/proto-1.0-SNAPSHOT.jar and b/proto/java/build/libs/proto-1.0-SNAPSHOT.jar differ`,
			expected: entity.FileDiffHeader{
				Status: entity.ChangeModified,
				FromFilePath: "proto/java/build/libs/proto-1.0-SNAPSHOT.jar",
				ToFilePath: "proto/java/build/libs/proto-1.0-SNAPSHOT.jar",
				Similarity: 0,
			},
			expectedHasErr: false,
		},
		{
			name: "image file modified",
			input: ` 
a/src/main/resources/static/img/default_avatar.jpg b/src/main/resources/static/img/default_avatar.jpeg
similarity index 100%
rename from src/main/resources/static/img/default_avatar.jpg
rename to src/main/resources/static/img/default_avatar.jpeg`,
			expected: entity.FileDiffHeader{
				Status: entity.ChangeRenamed,
				FromFilePath: "src/main/resources/static/img/default_avatar.jpg",
				ToFilePath: "src/main/resources/static/img/default_avatar.jpeg",
				Similarity: 100,
			},
			expectedHasErr: false,
		},
		{
			name: "invalid similarity",
			input: ` 
a/src/main/resources/static/img/default_avatar.jpg b/src/main/resources/static/img/default_avatar.jpeg
similarity index xx
rename from src/main/resources/static/img/default_avatar.jpg
rename to src/main/resources/static/img/default_avatar.jpeg`,
			expected: entity.FileDiffHeader{},
			expectedHasErr: true,
		},
		{
			name: "invalid hunk no filepath",
			input: ` 
a/src/main/resources/static/img/default_avatar.jpg b/src/main/resources/static/img/default_avatar.jpeg
similarity index 10`,
			expected: entity.FileDiffHeader{},
			expectedHasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := newFileDiffHeaderFromBlock(testCase.input)
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

func TestNewFileDiffFromBlock(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expected       entity.FileDiff
		expectedHasErr bool
	}{
		{
			name:           "empty block",
			input:          "",
			expected:       entity.FileDiff{},
			expectedHasErr: true,
		},
		{
			name: "file modified",
			input: `
 a/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java b/discussion/src/main/java/info/grouplive/discussion/service/CommentService.java
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
                 .stream()`,
			expected: entity.FileDiff{
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
							HeaderLine: "import info.grouplive.discussion.exceptions.PostNotFoundException;",
						},

						Lines: []entity.Line{
							{
								Status:  entity.LineUnchanged,
								Content: "import info.grouplive.discussion.mapper.CommentMapper;",
							},
							{
								Status:  entity.LineDeleted,
								Content: "import info.grouplive.discussion.model.User;",
							},
							{
								Status:  entity.LineDeleted,
								Content: "//import info.grouplive.discussion.model.User;",
							},
							{
								Status:  entity.LineAdded,
								Content: "import info.grouplive.discussion.model.UserModel;",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "import lombok.AllArgsConstructor;",
							},
						},
					},
					{
						HunkHeader: entity.HunkHeader{
							FromFileStartLine:  52,
							FromFileNumOfLines: 7,
							ToFileStartLine:    45,
							ToFileNumOfLines:   7,
							HeaderLine: "public class CommentService {",
						},

						Lines: []entity.Line{
							{
								Status:  entity.LineUnchanged,
								Content: "    }",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "    public List<CommentsDto> getAllCommentsForUser(String userName) {",
							},
							{
								Status:  entity.LineDeleted,
								Content: "        User user = userRepository.findByUsername(userName)",
							},
							{
								Status:  entity.LineAdded,
								Content: "        UserModel user = userRepository.findByUsername(userName)",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "                            .orElseThrow(() -> new UsernameNotFoundException(userName));",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "        return commentRepository.findAllByUser(user)",
							},
							{
								Status:  entity.LineUnchanged,
								Content: "                .stream()",
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewFileDiffFromBlock(testCase.input)
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

