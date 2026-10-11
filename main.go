package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/nonnika/ConventionalGit/model/cmdlineModel"
	"github.com/nonnika/ConventionalGit/model/commitModel"
	"github.com/nonnika/ConventionalGit/model/typesModel"
	"github.com/nonnika/ConventionalGit/tools/commitTools"
	"github.com/nonnika/ConventionalGit/tools/flagTools"
)

var version = "dev"

func main() {
	var options cmdlineModel.Options

	flag.BoolVar(&options.ShowVersion, "V", false, "显示构建版本号")
	flag.BoolVar(&options.ShowModifier, "M", false, "显示当前文档内容文件修改状况")
	flag.BoolVar(&options.DryRun, "dry-run", false, "不提交仅显示 Commit")
	flag.StringVar(&options.Scope, "scp", "", "代码修改适用的范围（可省略)")
	flag.StringVar(&options.Commit, "cmt", "", "代码提交的 Commit")
	flag.StringVar(&options.Numbers, "n", "", "需要 Add 的 FilePath 编号")

	// 注册所有短名称
	for _, tp := range typesModel.Types {
		flag.BoolFunc(tp.Short, tp.Name+" "+tp.Description, func(value string) error {
			if value == "false" {
				return nil
			}

			options.CommitType = tp.Name
			return nil
		})
	}

	flag.Usage = flagTools.PrintDefault
	flag.Parse()

	// 先按当前状态选择和暂存文件，再提交，避免提交改变文件编号。
	if options.Numbers != "" {
		if !runAdd(options) {
			return
		}
	}

	if options.CommitType != "" && options.Commit != "" {
		runCommit(options, options.DryRun)
	}

	if options.ShowVersion {
		fmt.Println(version)
	}

	if options.ShowModifier {
		err := commitTools.GetModifier()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func runCommit(options cmdlineModel.Options, dryRun bool) {
	cmt := commitModel.Commit{Tp: options.CommitType, Scope: options.Scope, Content: options.Commit}
	if dryRun {
		if err := flagTools.DryRun(cmt); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	out, err := commitTools.ExecCommit(cmt)
	fmt.Println(string(out))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runAdd(options cmdlineModel.Options) bool {
	fp, err := commitTools.GetFilePath(options.Numbers)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out, err := commitTools.ExecAdd(fp)
	if errors.Is(err, commitTools.ErrAddCanceled) {
		return false
	}
	fmt.Println(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return true
}
