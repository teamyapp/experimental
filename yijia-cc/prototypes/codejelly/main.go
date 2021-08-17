package main

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
)

func main() {
	repo := git.NewRepository("/Users/yijiacc/Documents/projects/laiprojects/grouplive")
	fileDiffs, err := repo.GetFileDiffsBetweenBranches("yijia-cc/feature-find-amenity-type", "master")
	if err != nil {
		panic(err)
	}

	for _, fileDiff := range fileDiffs {
		fmt.Println(fileDiff)
	}
}
