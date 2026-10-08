# ConventionalGit

一个用 Go 编写的命令行小工具，帮助你快速写出符合 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/) 规范的 Git 提交信息，同时提供带颜色高亮的文件修改状态查看。

## 特性

- 一条命令完成规范化提交，无需手动拼写 `type(scope): subject`
- 内置 11 种常用 commit type，每种都有单字母短参数
- 支持 scope（可选），自动拼接成 `type(scope): subject` 格式
- 彩色显示工作区修改状态，一眼区分已暂存 / 未暂存
- 按编号选择性暂存文件，确认后执行，避免误加
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

以下示例假设命令名为 `cg`。

### 提交代码

```bash
# 生成提交：feat:新增登录功能
cg -f -cmt "新增登录功能"

# 生成提交：fix(auth): 修复token过期判断错误
cg -x -scp auth -cmt "修复token过期判断错误"
```

参数说明：

| 参数 | 说明 |
| ---- | ---- |
| `-cmt` | 提交内容（与 type 参数配合使用） |
| `-scp` | 修改范围 scope，可省略 |
| `-M` | 查看当前文件修改状态 |
| `-n` | 按编号将对应文件加入暂存区（编号见 `-M` 输出） |
| `-V` | 显示构建版本号 |

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

每行开头的编号供 `-n` 选择文件使用。

### 暂存文件

先用 `-M` 查看编号，再用 `-n` 选择要暂存的文件，空格分隔、支持多个：

```bash
cg -n "0 2"
```

工具会先列出所选文件，输入 `Y`（或 `y` / `yes`）确认后才会执行 `git add`，输入其他内容则取消：

```
tools/x.go
a.txt
以上文件是您要添加的吗？输入 Y 以继续
```

编号超出范围或输入非法时会直接报错退出，不会执行任何 git 命令。

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
