package commitTools

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdlineModel"
	"github.com/nonnika/ConventionalGit/model/colorsModel"
	"github.com/nonnika/ConventionalGit/tools/gitTools"
)

func getModifier() ([]cmdlineModel.Status, error) {
	cmd := exec.Command("git", "-c", "core.quotePath=false", "status", "--porcelain", "-z")
	root, err := gitTools.GetRepoRoot()
	if err != nil {
		return nil, err
	}
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("读取 Git 状态失败: %w\n%s", err, out)
	}
	return parseStatus(out)
}

func parseStatus(out []byte) ([]cmdlineModel.Status, error) {
	if len(out) == 0 {
		return nil, nil
	}
	var stat []cmdlineModel.Status
	lines := strings.Split(strings.TrimSuffix(string(out), "\x00"), "\x00")
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) < 4 || lines[i][2] != ' ' {
			return nil, fmt.Errorf("无效的 Git 状态记录: %q", lines[i])
		}
		mod := lines[i][0:2]
		entry := cmdlineModel.Status{Mod: mod, FileNewPath: lines[i][3:], FileOldPath: lines[i][3:]}
		if strings.ContainsAny(mod, "RC") {
			i++
			if i >= len(lines) || lines[i] == "" {
				return nil, fmt.Errorf("Git 重命名或复制记录缺少原路径")
			}
			entry.FileOldPath = lines[i]
		}
		stat = append(stat, entry)
	}
	return stat, nil
}

func GetModifier() error {
	Stat, err := getModifier()
	if err != nil {
		return err
	}

	for index, stat := range Stat {
		col := colorsModel.RED
		chars := strings.Split(stat.Mod, "")
		if strings.ContainsAny(stat.Mod, "RC") {
			col = colorsModel.YELLOW
		}

		if chars[1] == " " {
			col = colorsModel.GREEN
		}
		fmt.Printf("%d %s%s%s", index, col, stat.Mod, colorsModel.END)

		if strings.ContainsAny(stat.Mod, "RC") {
			fmt.Printf(" %s -> %s\n", stat.FileOldPath, stat.FileNewPath)
			continue
		}

		fmt.Printf(" %s\n", stat.FileNewPath)

	}

	if len(Stat) == 0 {
		fmt.Printf("%s没有需要提交的内容%s\n", colorsModel.GREEN, colorsModel.END)
		return nil
	}
	fmt.Printf("合计：%s%d%s\n", colorsModel.YELLOW, len(Stat), colorsModel.END)
	return nil
}
