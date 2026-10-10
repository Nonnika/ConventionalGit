package commitTools

import (
	"fmt"
	"os/exec"

	"github.com/nonnika/ConventionalGit/model/commitModel"
)

// ExecCommit 执行命令
func ExecCommit(cmt commitModel.Commit) ([]byte, error) {
	return exec.Command("git", "commit", "-m", toString(cmt)).CombinedOutput()
}

func toString(cmt commitModel.Commit) string {
	if cmt.Scope == "" {
		return fmt.Sprintf("%s: %s", cmt.Tp, cmt.Content)
	}

	return fmt.Sprintf("%s(%s): %s", cmt.Tp, cmt.Scope, cmt.Content)
}
