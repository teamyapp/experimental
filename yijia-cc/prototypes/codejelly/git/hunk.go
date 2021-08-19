package git

import (
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"strings"
)

func newHunkFromLine(block string) (entity.Hunk, error) {
	parts := strings.Split(block, "@@")
	// Each
	//for _, part := range parts {
	//
	//}
}