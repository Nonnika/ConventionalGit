package flagTools

import (
	"fmt"

	"github.com/nonnika/ConventionalGit/model/colorsModel"
	"github.com/nonnika/ConventionalGit/model/commitModel"
	"github.com/nonnika/ConventionalGit/tools/commitTools"
)

func DryRun(cmt commitModel.Commit) error {
	files, err := commitTools.GetFilesStaged()
	if err != nil {
		return err
	}
	fmt.Printf("%s提交信息:%s %s\n", colorsModel.GREEN, colorsModel.END, commitTools.PraseCommit(cmt))
	if len(files) == 0 {
		fmt.Printf("%s没有任何文件被暂存%s\n", colorsModel.GREEN, colorsModel.END)
		return nil
	}
	fmt.Printf("\n%s暂存文件:%s\n", colorsModel.GREEN, colorsModel.END)
	for index, file := range files {
		fmt.Printf("\t%s%d%s %s\n", colorsModel.YELLOW, index, colorsModel.END, file)
	}
	fmt.Printf("合计: %s%d%s\n", colorsModel.YELLOW, len(files), colorsModel.END)
	return nil
}
