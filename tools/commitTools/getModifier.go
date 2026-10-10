package commitTools

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdlineModel"
	"github.com/nonnika/ConventionalGit/model/colorsModel"
	"github.com/nonnika/ConventionalGit/tools/git"
)

func getModifier() ([]cmdlineModel.Status, error) {
	var Stat []cmdlineModel.Status
	cmd := exec.Command("git", "-c", "core.quotePath=false", "status", "--porcelain")
	root, err := git.GetRepoRoot()
	if err != nil {
		return nil, err
	}
	cmd.Dir = root
	out, err := cmd.CombinedOutput()

	lines := strings.SplitSeq(strings.Trim(string(out), "\n"), "\n")

	for line := range lines {
		if line == "" {
			continue
		}

		mod := line[0:2]
		Stat = append(Stat, cmdlineModel.Status{Mod: mod, FilePath: line[3:]})
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
		if chars[1] == " " {
			col = colorsModel.GREEN
		}
		fmt.Printf("%d %s%s%s", index, col, stat.Mod, colorsModel.END)

		fmt.Println(" " + stat.FilePath)
	}

	if len(Stat) == 0 {
		fmt.Printf("%s没有需要提交的内容%s\n", colorsModel.GREEN, colorsModel.END)
		return nil
	}
	fmt.Printf("合计：%s%d%s\n", colorsModel.YELLOW, len(Stat), colorsModel.END)
	return nil
}
