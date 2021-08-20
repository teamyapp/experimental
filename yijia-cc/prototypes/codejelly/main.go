package main

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
)

func main() {
	git := git.NewGit("/Users/yijiacc/Documents/projects/laiprojects/grouplive")
	fileHeaders, err := git.GetFileDiffHeadersBetweenBranches("yijia-cc/feature-find-amenity-type", "master")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	for _, fileHeader := range fileHeaders {
		fmt.Println(fileHeader)
	}
	fmt.Println(fileHeaders)
}
