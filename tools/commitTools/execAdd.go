package commitTools

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/tools/git"
)

// ExecAdd 执行 Git add 命令
func ExecAdd(filePath []string) (string, error) {
	var confirm string
	for _, file := range filePath {
		fmt.Println(file)
	}
	fmt.Println("以上文件是您要添加的吗？输入 Y 以继续")
	_, err := fmt.Scanf("%s", &confirm)
	if err != nil {
		return "", err
	}
	confirm = strings.TrimSpace(confirm)
	confirm = strings.ToLower(confirm)
	if confirm != "y" && confirm != "yes" {
		fmt.Println("停止！！")
		return "", nil
	}

	args := append([]string{"add"}, filePath...)
	cmd := exec.Command("git", args...)

	root, err := git.GetRepoRoot()
	if err != nil {
		return "", err
	}

	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	return string(out), err
}
