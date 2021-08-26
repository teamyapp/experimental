package main

import (
	"fmt"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
)

func main() {
	g := git.NewRepo("/Users/yijiacc/Documents/projects/laiprojects/grouplive")
	fileDiffHeaders, err := g.GetFileDiffHeadersBetweenBranches("yijia-cc/feature-find-amenity-type", "master")

	if err != nil {
		panic(err)
	}

	fileDiff, err := g.GetFileDiffsBetweenBranches("yijia-cc/feature-find-amenity-type", "master")
	fmt.Println(len(fileDiffHeaders), len(fileDiff))
	fmt.Println(fileDiff[1].FileDiffHeader)

	hunks := fileDiff[11].Hunks
	for _, hunk := range hunks {
		lines := hunk.Lines
		fmt.Println(lines)
		//for _, _ := range lines {
		//	//fmt.Print("Line Status: ", line.Status)
		//	//fmt.Println(". Line Content: ", line.Content)
		//}
 	}
}
