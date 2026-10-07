package commit

import "fmt"

// PraseCommit 组合 Commit
func PraseCommit(tp, scope, content string) string {
	var commit string
	if scope != "" {
		commit = fmt.Sprintf("%s(%s): %s", tp, scope, content)
	} else {
		commit = fmt.Sprintf("%s: %s", tp, content)
	}
	return commit
}
