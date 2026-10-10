// Package commitModel
// 本包主要为 Commit 相关操作提供可行工具
package commitTools

import (
	"fmt"

	commit3 "github.com/nonnika/ConventionalGit/model/commitModel"
)

// PraseCommit 组合 Commit
func PraseCommit(cmt commit3.Commit) string {
	var commit string
	if cmt.Scope != "" {
		commit = fmt.Sprintf("%s(%s): %s", cmt.Tp, cmt.Scope, cmt.Content)
	} else {
		commit = fmt.Sprintf("%s: %s", cmt.Tp, cmt.Content)
	}
	return commit
}
