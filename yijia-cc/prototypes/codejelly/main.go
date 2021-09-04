package main

import (
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/service"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/stat"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/view"
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
	fullFileDiff, err := codeReview.GetFile(fileDiffs[9], "yijia-cc/feature-find-amenity-type")

	if err != nil {
		panic(err)
	}
/*
	for _, chunk := range fullFileDiff.Chunks {
		for _, line := range chunk.NumberedLines {

			fmt.Printf("[%02d, %02d]", line.FromFileLineNumber, line.ToFileLineNumber)
			if line.Status == entity.LineUnchanged {
				fmt.Print(" ")
			} else if line.Status == entity.LineDeleted {
				fmt.Print("-")
			} else if line.Status == entity.LineAdded {
				fmt.Print("+")
			}
			fmt.Println(line.Content)
		}
	}

 */
	stats, err := stat.ComputeFileDiffStats(fullFileDiff)
	if err != nil {
		panic(err)
	}
	fmt.Println(stats)

	//unifiedView := view.RenderUnifiedView(fullFileDiff)
	//
	//for _, chunk := range unifiedView.AllChunks {
	//	for _, line := range chunk.NumberedLines {
	//		fmt.Printf("[%02d, %02d]", line.FromFileLineNumber, line.ToFileLineNumber)
	//		if line.Status == entity.LineUnchanged {
	//			fmt.Print(" ")
	//		} else if line.Status == entity.LineDeleted {
	//			fmt.Print("-")
	//		} else if line.Status == entity.LineAdded {
	//			fmt.Print("+")
	//		}
	//		fmt.Println(line.Content)
	//	}
	//}
	//fmt.Println(unifiedView.HunkIndices)

	splitview, err := view.RenderSplitView(fullFileDiff)
	if err != nil {
		panic(err)
	}

	for _, chunkPair := range splitview.AllChunkPairs {
		for _, line := range chunkPair.ToFileChunk.NumberedLines {
			fmt.Printf("[%02d, %02d]", line.FromFileLineNumber, line.ToFileLineNumber)
			if line.Status == entity.LineUnchanged {
				fmt.Print(" ")
			} else if line.Status == entity.LineDeleted {
				fmt.Print("-")
			} else if line.Status == entity.LineAdded {
				fmt.Print("+")
			}
			fmt.Println(line.Content)
		}
	}
	fmt.Println(splitview.HunkPairIndices)
}
