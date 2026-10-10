// Package commitModel
// 本包定义了 Commit 的数据模型
package commitModel

type Commit struct {
	Tp      string
	Scope   string
	Content string
}
