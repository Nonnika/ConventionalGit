package commit

import (
	"fmt"
	"os/exec"

	"github.com/nonnika/ConventionalGit/model/commit"
)

// ExecCommit 执行命令
func ExecCommit(cmt commit.Commit) ([]byte, error) {
	return exec.Command("git", "commit", "-m", toString(cmt)).CombinedOutput()
}

func toString(cmt commit.Commit) string {
	if cmt.Scope == "" {
		return fmt.Sprintf("%s:%s", cmt.Tp, cmt.Content)
	}

	return fmt.Sprintf("%s(%s):%s", cmt.Tp, cmt.Scope, cmt.Content)
}
