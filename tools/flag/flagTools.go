// Package flag
// 为 Flag 相关操作提供自定义工具
package flag

import (
	"flag"
	"fmt"

	"github.com/nonnika/ConventionalGit/model/types"
)

func PrintDefault() {
	fmt.Println("用法：cg [options]")
	fmt.Printf("\nOptions:\n")

	flag.VisitAll(func(f *flag.Flag) {
		if types.IsType[f.Name] {
			return
		}
		fmt.Printf("\t-%s %s\n", f.Name, f.Usage)
	})

	fmt.Println("Commit Types:")
	for _, tp := range types.Types {
		fmt.Printf("\t-%s %s %s \n", tp.Short, tp.Name, tp.Description)
	}
}
