// Package flagTools
// 为 Flag 相关操作提供自定义工具
package flagTools

import (
	"flag"
	"fmt"

	"github.com/nonnika/ConventionalGit/model/typesModel"
)

// PrintDefault 打印所有的 Flag
func PrintDefault() {
	fmt.Println("Usage：\n\tcg [type] <options>")
	fmt.Println("Output: ")
	fmt.Println("\t<type>[optional scope]: <description>")
	fmt.Printf("\nOptions:\n")

	flag.VisitAll(func(f *flag.Flag) {
		if typesModel.IsType[f.Name] {
			return
		}
		fmt.Printf("\t-%s %s\n", f.Name, f.Usage)
	})

	fmt.Println("Commit Types:")
	for _, tp := range typesModel.Types {
		fmt.Printf("\t-%s %s %s \n", tp.Short, tp.Name, tp.Description)
	}
}
