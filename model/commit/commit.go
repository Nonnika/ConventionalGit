// Package commit
// 本包定义了 Commit 的数据模型
package commit

type Commit struct {
	Tp      string
	Scope   string
	Content string
}
