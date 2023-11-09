package main

import (
	"fmt"
	"github.com/qiuzhanghua/common/tz"
	"github.com/spf13/cobra"
)

var TzXCmd = &cobra.Command{
	Use:   "x",
	Short: "Extract zip",
	Long:  "Extract files from a zip",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tz x <zip> [outDir]")
			return
		}
		zipName := args[0]
		var outDir string
		if len(args) < 2 {
			outDir = "."
		} else {
			outDir = args[1]
		}
		tz.Extract(zipName, outDir)
	},
}
