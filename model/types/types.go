// Package types
// 本包声明了常见的 commit types
package types

type CommitType struct {
	Short       string
	Name        string
	Description string
}

var Types = []CommitType{
	{"f", "feat", "新增功能"},
	{"x", "fix", "修复缺陷"},
	{"d", "docs", "文档变更"},
	{"s", "style", "代码格式调整"},
	{"r", "refactor", "代码重构"},
	{"p", "perf", "性能优化"},
	{"t", "test", "新增或更新测试"},
	{"b", "build", "构建系统变更"},
	{"c", "ci", "CI 配置变更"},
	{"o", "chore", "其他杂项维护"},
	{"v", "revert", "回退之前的提交"},
}

var IsType = map[string]bool{
	"f": true,
	"x": true,
	"d": true,
	"s": true,
	"r": true,
	"p": true,
	"t": true,
	"b": true,
	"c": true,
	"o": true,
	"v": true,
}
