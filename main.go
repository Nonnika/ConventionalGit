package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nonnika/ConventionalGit/model/cmdline"
	commit3 "github.com/nonnika/ConventionalGit/model/commit"
	"github.com/nonnika/ConventionalGit/model/types"
	commit2 "github.com/nonnika/ConventionalGit/tools/commit"
	flag2 "github.com/nonnika/ConventionalGit/tools/flag"
)

var version = "dev"

func main() {
	var options cmdline.Options

	flag.BoolVar(&options.ShowVersion, "V", false, "显示构建版本号")
	flag.BoolVar(&options.ShowModifier, "M", false, "显示当前文档内容文件修改状况")
	flag.StringVar(&options.Scope, "scp", "", "代码修改适用的范围（可省略)")
	flag.StringVar(&options.Commit, "cmt", "", "代码提交的 Commit")
	flag.StringVar(&options.Numbers, "n", "", "需要 Add 的 FilePath 编号")

	// 注册所有短名称
	for _, tp := range types.Types {
		flag.BoolFunc(tp.Short, tp.Name+" "+tp.Description, func(value string) error {
			if value == "false" {
				return nil
			}

			options.CommitType = tp.Name
			return nil
		})
	}

	flag.Usage = flag2.PrintDefault
	flag.Parse()

	// 处理提交 Commit 方法
	if options.CommitType != "" && options.Commit != "" {
		runCommit(options)
	}

	// 处理 Add 方法
	if options.Numbers != "" {
		runAdd(options)
	}

	if options.ShowVersion {
		fmt.Println(version)
	}

	if options.ShowModifier {
		err := commit2.GetModifier()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func runCommit(options cmdline.Options) {
	cmt := commit3.Commit{Tp: options.CommitType, Scope: options.Scope, Content: options.Commit}
	out, err := commit2.ExecCommit(cmt)
	fmt.Println(string(out))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runAdd(options cmdline.Options) {
	fp, err := commit2.GetFilePath(options.Numbers)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out, err := commit2.ExecAdd(fp)
	fmt.Println(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
