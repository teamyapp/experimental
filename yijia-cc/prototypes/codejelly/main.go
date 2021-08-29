package main

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/service"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
)

func main() {
	g := git.NewRepo("/Users/yijiacc/Documents/projects/laiprojects/grouplive")
	fileDiffHeaders, err := g.GetFileDiffHeadersBetweenBranches("yijia-cc/feature-find-amenity-type", "master")

	if err != nil {
		panic(err)
	}

	fileDiffs, err := g.GetFileDiffsBetweenBranches("yijia-cc/feature-find-amenity-type", "master")
	fmt.Println(len(fileDiffHeaders), len(fileDiffs))
	fmt.Println(fileDiffs[1].FileDiffHeader)

	hunks := fileDiffs[11].Hunks
	for _, hunk := range hunks {
		lines := hunk.Lines
		fmt.Println(len(lines))
		//for _, _ := range lines {
		//	//fmt.Print("Line Status: ", line.Status)
		//	//fmt.Println(". Line Content: ", line.Content)
		//}
 	}

	codeReview := service.NewCodeReview(g)
	fullFileDiff, err := codeReview.GetFile(fileDiffs[0], "yijia-cc/feature-find-amenity-type")
	if err != nil {
		panic(err)
	}

	for _, chunk := range fullFileDiff.Chunks {
		for _, line := range chunk.Lines {
			fmt.Println(line.Content)
		}
	}
}
