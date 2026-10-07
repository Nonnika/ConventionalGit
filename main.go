package main

import (
	"flag"
	"fmt"
)

var version = "dev"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "显示构建版本号")
	flg := flag.String("op", "", "操作类型")
	flag.Usage = func() {
		fmt.Println("用法：cg [options]")
		fmt.Printf("\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*flg)
}
