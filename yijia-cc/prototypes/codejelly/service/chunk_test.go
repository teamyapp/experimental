package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"testing"
)

func TestGetChunks(t *testing.T) {
	testCases := []struct {
		name string
		block string
		fromFile string
		expected []entity.Chunk
	} {
		{
			name: "multiple hunks",
			block:
`diff --git a/calendar/gql/resolver/scalar.go b/calendar/gql/resolver/scalar.go
index 7ac4bc5..6aaa9e6 100644
--- a/calendar/gql/resolver/scalar.go
+++ b/calendar/gql/resolver/scalar.go
@@ -9,7 +9,7 @@ import (
 var _ decode.Unmarshaler = (*Duration)(nil)
 
 type Duration struct {
-	time.Duration
+	duration time.Duration
 }
 
 func (d Duration) ImplementsGraphQLType(name string) bool {
@@ -19,9 +19,9 @@ func (d Duration) ImplementsGraphQLType(name string) bool {
 func (d *Duration) UnmarshalGraphQL(input interface{}) error {
 	switch input := input.(type) {
 	case time.Duration:
-		d.Duration = input
+		d.duration = input
 	case int:
-		d.Duration = time.Duration(input)
+		d.duration = time.Duration(input)
 	}
 	return nil
 }`,
			fromFile:
`package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go/decode"
)

var _ decode.Unmarshaler = (*Duration)(nil)

type Duration struct {
	time.Duration
}

func (d Duration) ImplementsGraphQLType(name string) bool {
	return name == "Duration"
}

func (d *Duration) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Duration:
		d.Duration = input
	case int:
		d.Duration = time.Duration(input)
	}
	return nil
}`,
			expected: []entity.Chunk {
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status:  entity.LineUnchanged,
							Content: "package resolver",
							FromFileLineNumber: 1,
							ToFileLineNumber: 1,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 2,
							ToFileLineNumber: 2,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "import (",
							FromFileLineNumber: 3,
							ToFileLineNumber: 3,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "\t\"time\"",
							FromFileLineNumber: 4,
							ToFileLineNumber: 4,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 5,
							ToFileLineNumber: 5,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "\t\"github.com/graph-gophers/graphql-go/decode\"",
							FromFileLineNumber: 6,
							ToFileLineNumber: 6,
						},
						{
							Status:  entity.LineUnchanged,
							Content: ")",
							FromFileLineNumber: 7,
							ToFileLineNumber: 7,
						},
						{
							Status:  entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 8,
							ToFileLineNumber: 8,
						},
					},
					IsHunk: false,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "var _ decode.Unmarshaler = (*Duration)(nil)",
							FromFileLineNumber: 9,
							ToFileLineNumber: 9,
						},
						{
							Status: entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 10,
							ToFileLineNumber: 10,
						},
						{
							Status: entity.LineUnchanged,
							Content: "type Duration struct {",
							FromFileLineNumber: 11,
							ToFileLineNumber: 11,
						},
						{
							Status: entity.LineDeleted,
							Content: "\ttime.Duration",
							FromFileLineNumber: 12,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineAdded,
							Content: "\tduration time.Duration",
							FromFileLineNumber: entity.NoLineNumber,
							ToFileLineNumber: 12,
						},
						{
							Status: entity.LineUnchanged,
							Content: "}",
							FromFileLineNumber: 13,
							ToFileLineNumber: 13,
						},
						{
							Status: entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 14,
							ToFileLineNumber: 14,
						},
						{
							Status: entity.LineUnchanged,
							Content: "func (d Duration) ImplementsGraphQLType(name string) bool {",
							FromFileLineNumber: 15,
							ToFileLineNumber: 15,
						},
					},
					IsHunk: true,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "\treturn name == \"Duration\"",
							FromFileLineNumber: 16,
							ToFileLineNumber: 16,
						},
						{
							Status: entity.LineUnchanged,
							Content: "}",
							FromFileLineNumber: 17,
							ToFileLineNumber: 17,
						},
						{
							Status: entity.LineUnchanged,
							Content: "",
							FromFileLineNumber: 18,
							ToFileLineNumber: 18,
						},
					},
					IsHunk: false,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "func (d *Duration) UnmarshalGraphQL(input interface{}) error {",
							FromFileLineNumber: 19,
							ToFileLineNumber: 19,
						},
						{
							Status: entity.LineUnchanged,
							Content: "\tswitch input := input.(type) {",
							FromFileLineNumber: 20,
							ToFileLineNumber: 20,
						},
						{
							Status: entity.LineUnchanged,
							Content: "\tcase time.Duration:",
							FromFileLineNumber: 21,
							ToFileLineNumber: 21,
						},
						{
							Status: entity.LineDeleted,
							Content: "\t\td.Duration = input",
							FromFileLineNumber: 22,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineAdded,
							Content: "\t\td.duration = input",
							FromFileLineNumber: entity.NoLineNumber,
							ToFileLineNumber: 22,
						},
						{
							Status: entity.LineUnchanged,
							Content: "\tcase int:",
							FromFileLineNumber: 23,
							ToFileLineNumber: 23,
						},
						{
							Status: entity.LineDeleted,
							Content: "\t\td.Duration = time.Duration(input)",
							FromFileLineNumber: 24,
							ToFileLineNumber: entity.NoLineNumber,
						},
						{
							Status: entity.LineAdded,
							Content: "\t\td.duration = time.Duration(input)",
							FromFileLineNumber: entity.NoLineNumber,
							ToFileLineNumber: 24,
						},
						{
							Status: entity.LineUnchanged,
							Content: "\t}",
							FromFileLineNumber: 25,
							ToFileLineNumber: 25,
						},
						{
							Status: entity.LineUnchanged,
							Content: "\treturn nil",
							FromFileLineNumber: 26,
							ToFileLineNumber: 26,
						},
						{
							Status: entity.LineUnchanged,
							Content: "}",
							FromFileLineNumber: 27,
							ToFileLineNumber: 27,
						},
					},
					IsHunk: true,
				},
			},
		},
		{
			name: "added hunk",
			block:
`diff --git a/dashboard/auth/authenticator.go b/dashboard/auth/authenticator.go
new file mode 100644
index 0000000..0511e93
--- /dev/null
+++ b/dashboard/auth/authenticator.go
@@ -0,0 +1,5 @@
+package auth
+
+type Authenticator interface {
+	VerifyIdentity(authToken string) (string, error)
+}`,
			fromFile: ``,
			expected: []entity.Chunk{
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineAdded,
							Content: "package auth",
							FromFileLineNumber: -1,
							ToFileLineNumber: 1,
						},
						{
							Status: entity.LineAdded,
							Content: "",
							FromFileLineNumber: -1,
							ToFileLineNumber: 2,
						},
						{
							Status: entity.LineAdded,
							Content: "type Authenticator interface {",
							FromFileLineNumber: -1,
							ToFileLineNumber: 3,
						},
						{
							Status: entity.LineAdded,
							Content: "\tVerifyIdentity(authToken string) (string, error)",
							FromFileLineNumber: -1,
							ToFileLineNumber: 4,
						},
						{
							Status: entity.LineAdded,
							Content: "}",
							FromFileLineNumber: -1,
							ToFileLineNumber: 5,
						},
					},
					IsHunk: true,
				},
			},
		},
		{
			name: "non empty fromFile with one hunk",
			block:
`diff --git a/dashboard/entity/id.go b/dashboard/entity/id.go
index 0000000..acf3183 100644
--- a/dashboard/entity/id.go
+++ b/dashboard/entity/id.go
@@ -2,3 +2,3 @@
+package entity
+
+type ID string`,
			fromFile:
`package dao`,
			expected: []entity.Chunk{
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineUnchanged,
							Content: "package dao",
							FromFileLineNumber: 1,
							ToFileLineNumber: 1,
						},
					},
					IsHunk: false,
				},
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineAdded,
							Content: "package entity",
							FromFileLineNumber: -1,
							ToFileLineNumber: 2,
						},
						{
							Status: entity.LineAdded,
							Content: "",
							FromFileLineNumber: -1,
							ToFileLineNumber: 3,
						},
						{
							Status: entity.LineAdded,
							Content: "type ID string",
							FromFileLineNumber: -1,
							ToFileLineNumber: 4,
						},
					},
					IsHunk: true,
				},
			},
		},
		{
			name: "deleted block",
			block:
`diff --git a/dashboard/entity/id.go b/dashboard/entity/id.go
index 0000000..acf3183 100644
--- a/dashboard/entity/id.go
+++ /dev/null
@@ -1,4 +0,0 @@
-package dao
-package entity
-
-type ID string`,
			fromFile:
`package dao
package entity

type ID string`,
			expected: []entity.Chunk{
				{
					NumberedLines: []entity.NumberedLine{
						{
							Status: entity.LineDeleted,
							Content: "package dao",
							FromFileLineNumber: 1,
							ToFileLineNumber: -1,
						},
						{
							Status: entity.LineDeleted,
							Content: "package entity",
							FromFileLineNumber: 2,
							ToFileLineNumber: -1,
						},
						{
							Status: entity.LineDeleted,
							Content: "",
							FromFileLineNumber: 3,
							ToFileLineNumber: -1,
						},
						{
							Status: entity.LineDeleted,
							Content: "type ID string",
							FromFileLineNumber: 4,
							ToFileLineNumber: -1,
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
			hunks, _, err := git.NewHunksFromBlock(testCase.block)
			if err != nil {
				assert.Fail(t, "fail")
			}
			actual := GetChunks(hunks, testCase.fromFile)

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
