package commit

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdline"
	"github.com/nonnika/ConventionalGit/model/colors"
	"github.com/nonnika/ConventionalGit/tools/git"
)

func getModifier() ([]cmdline.Status, error) {
	var Stat []cmdline.Status
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
		Stat = append(Stat, cmdline.Status{Mod: mod, FilePath: line[3:]})
	}

	return Stat, err
}

func GetModifier() error {
	Stat, err := getModifier()
	if err != nil {
		return err
	}

	for index, stat := range Stat {
		col := colors.RED
		chars := strings.Split(stat.Mod, "")
		if chars[1] == " " {
			col = colors.GREEN
		}
		fmt.Printf("%d %s%s%s", index, col, stat.Mod, colors.END)

		fmt.Println(" " + stat.FilePath)
	}

	fmt.Printf("合计：%s%d%s\n", colors.YELLOW, len(Stat), colors.END)
	return nil
}
