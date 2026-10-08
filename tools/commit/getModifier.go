package commit

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdline"
	"github.com/nonnika/ConventionalGit/model/colors"
)

func getModifier() ([]cmdline.Status, error) {
	var Stat []cmdline.Status

	out, err := exec.Command("git", "status", "--porcelain").Output()

	lines := strings.SplitSeq(strings.Trim(string(out), "\n"), "\n")

	for line := range lines {
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

	for _, stat := range Stat {
		col := colors.RED
		chars := strings.Split(stat.Mod, "")
		if chars[1] == " " {
			col = colors.GREEN
		}
		fmt.Printf("%s%s%s", col, stat.Mod, colors.END)

		fmt.Println(" " + stat.FilePath)
	}

	fmt.Printf("合计：%s%d%s\n", colors.YELLOW, len(Stat), colors.END)
	return nil
}
