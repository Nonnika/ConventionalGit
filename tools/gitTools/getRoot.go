// Package gitTools
// 本包为 Git 提供辅助工具
package gitTools

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetRepoRoot 返回 Git 基础工作目录
func GetRepoRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("获取 Git 仓库根目录失败: %w", err)
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}
