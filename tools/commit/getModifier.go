package commit

import (
	"os/exec"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdline"
)

// GetModifier 获取 Git status 输出
func GetModifier() ([]cmdline.Status, error) {
	var Stat []cmdline.Status

	out, err := exec.Command("git", "status", "-s").Output()

	lines := strings.SplitSeq(strings.Trim(string(out), "\n"), "\n")

	for line := range lines {
		words := strings.Fields(line)
		Stat = append(Stat, cmdline.Status{Mod: words[0], FilePath: words[1]})
	}

	return Stat, err
}
