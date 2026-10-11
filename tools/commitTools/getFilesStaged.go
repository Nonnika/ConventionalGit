package commitTools

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/tools/gitTools"
)

// GetFilesStaged 获取暂存的文件
func GetFilesStaged() ([]string, error) {
	root, err := gitTools.GetRepoRoot()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("git", "diff", "--staged", "--name-only", "-z")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("读取暂存文件失败: %w\n%s", err, out)
	}
	if len(out) == 0 {
		return nil, nil
	}
	files := strings.Split(strings.TrimSuffix(string(out), "\x00"), "\x00")
	return files, nil
}
