package commit

import (
	"fmt"
	"os/exec"
	"strings"
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
	out, err := exec.Command("git", args...).CombinedOutput()
	return string(out), err
}
