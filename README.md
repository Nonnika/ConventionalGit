# ConventionalGit

一个用 Go 编写的命令行小工具，帮助你快速写出符合 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/) 规范的 Git 提交信息，同时提供带颜色高亮的文件修改状态查看。

## 特性

- 一条命令完成规范化提交，无需手动拼写 `type(scope): subject`
- 内置 11 种常用 commit type，每种都有单字母短参数
- 支持 scope（可选），自动拼接成 `type(scope): subject` 格式
- 彩色显示工作区修改状态，一眼区分已暂存 / 未暂存
- 按单个编号、多个编号或连续范围选择文件，确认后暂存
- 使用 `--dry-run` 预览提交信息和当前已暂存文件
- 支持在仓库子目录中查看状态和暂存文件
- 构建时注入版本号，方便追溯

## 安装

从源码构建（需要 Go 环境）：

```bash
git clone https://github.com/nonnika/ConventionalGit.git
cd ConventionalGit
make build
```

构建产物位于 `build/ConventionalGit`。建议建立软链接为 `cg` 使用：

```bash
ln -s "$(pwd)/build/ConventionalGit" /usr/local/bin/cg
```

也可以直接安装：

```bash
go install github.com/nonnika/ConventionalGit@latest
```

## 使用

以下示例假设命令名为 `cg`，在 Git 仓库内执行。

### 提交代码

```bash
# 提交当前暂存区，生成提交信息：feat: 新增登录功能
cg -f -cmt "新增登录功能"

# 生成提交：fix(auth): 修复token过期判断错误
cg -x -scp auth -cmt "修复token过期判断错误"
```

提交前需先暂存要提交的改动，可以使用 `git add` 或下面的 `cg -n`。提交命令会提交当前暂存区中的全部改动。

参数说明：

| 参数 | 说明 |
| ---- | ---- |
| `-cmt` | 提交内容（与 type 参数配合使用） |
| `-scp` | 修改范围 scope，可省略 |
| `-M` | 查看当前文件修改状态 |
| `-n` | 按编号或连续范围将对应文件加入暂存区（编号见 `-M` 输出） |
| `--dry-run` | 与 type 和 `-cmt` 配合，预览提交信息及已暂存文件，不执行提交 |
| `-V` | 显示构建版本号 |
| `-h` | 显示帮助和全部提交类型 |

### 查看修改状态

```bash
cg -M
```

输出示例：

```
0  M tools/x.go   # 红色：存在未暂存的修改
1 M  main.go      # 绿色：修改已全部暂存
2 ?? a.txt        # 红色：未跟踪文件
合计：3            # 黄色：变更文件总数
```

每行开头的编号供 `-n` 选择文件使用，从 `0` 开始。编号来自当前状态输出，文件状态变化后应重新查看。

没有变更时会显示“没有需要提交的内容”。在仓库子目录中执行时，也会显示整个仓库的变更，路径以仓库根目录为基准。

### 暂存文件

先用 `-M` 查看编号，再用 `-n` 选择要暂存的文件：

```bash
# 暂存编号为 0 的文件
cg -n "0"

# 暂存编号为 0 和 2 的文件，多个编号用空格分隔
cg -n "0 2"

# 暂存编号为 0 到 2 的文件，包含两端，共 3 项
cg -n "0-2"
```

连续范围的格式为 `起始编号-结束编号`，结束编号必须大于或等于起始编号。目前一次可使用多个单独编号或一个连续范围，不能混写成 `"0 2-4"`。

工具会先列出所选文件，输入 `Y`（或 `y` / `yes`）确认后才会执行 `git add`，输入其他内容则取消：

```
tools/x.go
a.txt
以上文件是您要添加的吗？输入 Y 以继续
```

编号超出范围、范围顺序错误或输入非法时会报错退出，不执行 `git add`。

### 预览提交

在提交参数后加上 `--dry-run`，查看生成的提交信息和当前已暂存文件：

```bash
cg -f -scp auth -cmt "新增登录功能" --dry-run
```

输出示例（省略颜色）：

```text
提交信息: feat(auth): 新增登录功能

暂存文件:
    0 main.go
    1 tools/auth.go
合计: 2
```

预览列表中的编号只用于展示；选择暂存文件时仍以 `cg -M` 的编号为准。

`--dry-run` 只跳过 `git commit`，不影响 `-n` 的暂存操作。建议分步暂存、预览和提交：

```bash
cg -M
cg -n "0-2"
cg -f -scp auth -cmt "新增登录功能" --dry-run
cg -f -scp auth -cmt "新增登录功能"
```

### 查看帮助

```bash
cg -h
```

## Commit Types

| 短参数 | 类型 | 说明 |
| ------ | ---- | ---- |
| `-f` | feat | 新增功能 |
| `-x` | fix | 修复缺陷 |
| `-d` | docs | 文档变更 |
| `-s` | style | 代码格式调整 |
| `-r` | refactor | 代码重构 |
| `-p` | perf | 性能优化 |
| `-t` | test | 新增或更新测试 |
| `-b` | build | 构建系统变更 |
| `-c` | ci | CI 配置变更 |
| `-o` | chore | 其他杂项维护 |
| `-v` | revert | 回退之前的提交 |

## 开发

```bash
make run-test   # 运行全部测试
make build      # 构建到 build/，版本号取自 git describe
```

## License

[Apache-2.0](LICENSE)
