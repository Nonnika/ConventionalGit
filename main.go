package main

import (
	"flag"
	"fmt"
	"os"

	commit3 "github.com/nonnika/ConventionalGit/model/commit"
	"github.com/nonnika/ConventionalGit/model/types"
	commit2 "github.com/nonnika/ConventionalGit/tools/commit"
	flag2 "github.com/nonnika/ConventionalGit/tools/flag"
)

var version = "dev"

func main() {
	var showVersion bool
	var commitType string
	var showModifier bool
	flag.BoolVar(&showVersion, "V", false, "显示构建版本号")
	flag.BoolVar(&showModifier, "M", false, "显示当前文档内容文件修改状况")
	scope := flag.String("scp", "", "代码修改适用的范围（可省略)")
	commit := flag.String("cmt", "", "代码提交的 Commit")
	numbers := flag.String("n", "", "需要 Add 的 FilePath 编号")

	// 注册所有短名称
	for _, tp := range types.Types {
		flag.BoolFunc(tp.Short, tp.Name+" "+tp.Description, func(value string) error {
			if value == "false" {
				return nil
			}

			commitType = tp.Name
			return nil
		})
	}

	flag.Usage = flag2.PrintDefault
	flag.Parse()

	if commitType != "" && *commit != "" {
		cmt := commit3.Commit{Tp: commitType, Scope: *scope, Content: *commit}
		out, err := commit2.ExecCommit(cmt)
		fmt.Println(string(out))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if numbers != nil && *numbers != "" {
		fp, err := commit2.GetFilePath(*numbers)
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

	if showVersion {
		fmt.Println(version)
	}

	if showModifier {
		err := commit2.GetModifier()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
