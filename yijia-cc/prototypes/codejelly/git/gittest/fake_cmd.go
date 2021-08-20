package gittest

import (
	"errors"
	"strings"

	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
)

var _ git.CommandExecutor = (*StubCommandExecutor)(nil)

type StubCommandExecutor struct {
	outputMap map[string]string
}

func (s StubCommandExecutor) Execute(cmd string, args ...string) (string, error){
	// get fromBranch and toBranch names from args
	branchArgs := args[len(args)-2:]
	output, ok := s.outputMap[strings.Join(branchArgs, ",")]
	if !ok {
		return "", errors.New("branch does not exists")
	}
	return output, nil
}

func NewStubCommandExecutor(outputMap map[string]string) StubCommandExecutor{
	return StubCommandExecutor {
		outputMap: outputMap,
	}
}
