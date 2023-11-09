package main

import (
	"fmt"
	"github.com/qiuzhanghua/common/tz"
	"github.com/spf13/cobra"
)

var TzCCmd = &cobra.Command{
	Use:   "c",
	Short: "Compress zip",
	Long:  "Compress files into a zip",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tz c <tar.gz> [dir|file ...]")
			return
		}
		zipName := args[0]
		tz.Compress(zipName, args[1:]...)
	},
}
