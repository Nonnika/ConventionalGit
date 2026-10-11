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
	var Stat []cmdlineModel.Status
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
	if len(out) == 0 {
		return nil, nil
	}

	lines := strings.Split(strings.Trim(string(out), "\x00"), "\x00")

	rename := false
	n := len(lines)
	for i := range n {
		if rename {
			Stat[len(Stat)-1].FileOldPath = lines[i]
			continue
		}
		if len(lines[i]) < 4 || lines[i][2] != ' ' {
			return nil, fmt.Errorf("无效的 Git 状态记录: %q", lines[i])
		}
		mod := lines[i][0:2]
		if string(mod[0]) == "R" {
			rename = true
		}

		Stat = append(Stat, cmdlineModel.Status{Mod: mod, FileNewPath: lines[i][3:], FileOldPath: lines[i][3:]})

	}

	return Stat, err
}

func GetModifier() error {
	Stat, err := getModifier()
	if err != nil {
		return err
	}

	for index, stat := range Stat {
		col := colorsModel.RED
		chars := strings.Split(stat.Mod, "")
		if chars[0] == "R" {
			col = colorsModel.YELLOW
		}

		if chars[1] == " " {
			col = colorsModel.GREEN
		}
		fmt.Printf("%d %s%s%s", index, col, stat.Mod, colorsModel.END)

		if chars[0] == "R" {
			fmt.Printf(" %s -> %s\n", stat.FileOldPath, stat.FileNewPath)
			return nil
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
