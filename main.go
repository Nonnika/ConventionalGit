package main

import (
	"flag"
	"fmt"

	"github.com/nonnika/ConventionalGit/model/types"
	flag2 "github.com/nonnika/ConventionalGit/tools/flag"
)

var version = "dev"

func main() {
	var showVersion bool
	var commitType string
	flag.BoolVar(&showVersion, "V", false, "显示构建版本号")
	scope := flag.String("scp", "", "代码修改适用的范围（可省略)")
	commit := flag.String("cmt", "", "代码提交的 Commit")

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

	if commitType != "" {
		message := PraseCommit(commitType, *scope, *commit)
		fmt.Println(message)
	}
}

// PraseCommit 组合 Commit
func PraseCommit(tp, scope, content string) string {
	var commit string
	if scope != "" {
		commit = fmt.Sprintf("%s(%s): %s", tp, scope, content)
	} else {
		commit = fmt.Sprintf("%s: %s", tp, content)
	}
	return commit
}
