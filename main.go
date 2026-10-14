package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nonnika/ConventionalGit/model/cmdlineModel"
	"github.com/nonnika/ConventionalGit/model/colorsModel"
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
			enabled, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			if !enabled {
				return nil
			}
			if options.CommitType != "" && options.CommitType != tp.Name {
				return fmt.Errorf("只能指定一个提交类型")
			}

			options.CommitType = tp.Name
			return nil
		})
	}

	flag.Usage = flagTools.PrintDefault
	flag.Parse()
	provided := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { provided[f.Name] = true })
	if err := validateOptions(options, flag.Args(), provided); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

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
		out := fmt.Sprintf("Version: %s\nMotto:%s Will You Will%s", version, colorsModel.GREEN, colorsModel.END)
		fmt.Println(out)
	}

	if options.ShowModifier {
		err := commitTools.GetModifier()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func validateOptions(options cmdlineModel.Options, args []string, provided map[string]bool) error {
	if len(args) != 0 {
		return fmt.Errorf("无法识别的位置参数: %s", strings.Join(args, " "))
	}
	if provided["cmt"] && strings.TrimSpace(options.Commit) == "" {
		return fmt.Errorf("-cmt 提交内容不能为空")
	}
	if provided["n"] && strings.TrimSpace(options.Numbers) == "" {
		return fmt.Errorf("-n 文件编号不能为空")
	}
	hasType := options.CommitType != ""
	hasContent := strings.TrimSpace(options.Commit) != ""
	if hasType != hasContent {
		return fmt.Errorf("提交类型与 -cmt 提交内容必须同时提供")
	}
	if (options.DryRun || options.Scope != "") && !hasType {
		return fmt.Errorf("--dry-run 和 -scp 需要同时提供提交类型与 -cmt 提交内容")
	}
	return nil
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
