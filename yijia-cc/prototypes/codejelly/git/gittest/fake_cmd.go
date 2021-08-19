package gittest

import (
	"errors"
	"fmt"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/git"
	"strings"
)

var _ git.CommandExecutor = (*StubCommandExecutor)(nil)

type StubCommandExecutor struct {
	outputMap map[string]string
}

func (s StubCommandExecutor) Execute(cmd string, args ...string) (string, error){
	fmt.Println(strings.Join(args[len(args)-2:], ","))
	output, ok := s.outputMap[strings.Join(args[len(args)-2:], ",")]
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

var outputs = []string {
		"A       dashboard/model/event.go\n" +
		"D       dashboard/seeds/seeder.go\n" +
		"M       discussion/.gitignore\n" +
		"R058    discussion/src/main/java/info/UserModel.java discussion/src/main/java/info/User.java",
}