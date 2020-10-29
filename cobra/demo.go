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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
