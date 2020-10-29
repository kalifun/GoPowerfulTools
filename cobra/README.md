![](https://cloud.githubusercontent.com/assets/173412/10886352/ad566232-814f-11e5-9cd0-aa101788c117.png)
# CoBra
> Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.
## 一、概念
> Cobra是建立在命令、参数和标志的结构上的。

Commands 代表操作, Args 代表事件 and Flags 操作的修饰符.
## 二、示例
```bash
  gh
Work seamlessly with GitHub from the command line.

USAGE
  gh <command> <subcommand> [flags]

CORE COMMANDS
  gist:       Create gists
  issue:      Manage issues
  pr:         Manage pull requests
  release:    Manage GitHub releases
  repo:       Create, clone, fork, and view repositories

ADDITIONAL COMMANDS
  alias:      Create command shortcuts
  api:        Make an authenticated GitHub API request
  auth:       Login, logout, and refresh your authentication
  completion: Generate shell completion scripts
  config:     Manage configuration for gh
  help:       Help about any command

FLAGS
  --help      Show help for command
  --version   Show gh version

EXAMPLES
  $ gh issue create
  $ gh repo clone cli/cli
  $ gh pr checkout 321

ENVIRONMENT VARIABLES
  See 'gh help environment' for the list of supported environment variables.

LEARN MORE
  Use 'gh <command> <subcommand> --help' for more information about a command.
  Read the manual at https://cli.github.com/manual

FEEDBACK
  Open an issue using 'gh issue create -R cli/cli'
```
上面是GitHub的CLI命令行工具。下面我们将会利用`Cobra`实现一个命令行工具。
## 三、Demo
### 下载
```
go get github.com/spf13/cobra
```
### 创建一个操作
``` golang
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Short是“help”输出中显示的简短描述。",
	Long:  `Long是“help<this command>”输出中显示的长消息。`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) < 1 {
			tip()
		}
	},
}

func tip() {
	fmt.Println("你可以尝试使用-h 来查询更多使用方法")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```
### 效果
```bash
  go run demo.go 
你可以尝试使用-h 来查询更多使用方法
```
```bash
  go run demo.go -h
Long是“help<this command>”输出中显示的长消息。

Usage:
  cobra [flags]

Flags:
  -h, --help   help for cobra
```
### 优化
```golang
var rootCmd1 = &cobra.Command{
	Use:   "cobra1",
	Short: "Short是“help”输出中显示的简短描述。",
	Long:  `Long是“help<this command>”输出中显示的长消息。`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) < 1 {
			tip()
		}
	},
}

func init() {
	rootCmd.AddCommand(rootCmd1)
}
```
```bash
  go run demo.go -h
Long是“help<this command>”输出中显示的长消息。

Usage:
  cobra [flags]
  cobra [command]

Available Commands:
  cobra1      Short是“help”输出中显示的简短描述。
  help        Help about any command

Flags:
  -h, --help   help for cobra

Use "cobra [command] --help" for more information about a command.
```