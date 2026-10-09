package commit

import (
	"os/exec"
	"strings"
)

// GetFilesStaged 获取暂存的文件
func GetFilesStaged() []string {
	out, _ := exec.Command("git", "diff", "--staged", "--name-only", "-z").Output()
	files := strings.Split(strings.TrimRight(string(out), "\x00"), "\x00")
	return files
}
