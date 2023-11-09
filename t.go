package main

import (
	"fmt"
	"github.com/qiuzhanghua/common/tz"
	"github.com/spf13/cobra"
)

var TzTCmd = &cobra.Command{
	Use:   "t",
	Short: "List zip",
	Long:  "List files from a zip",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			fmt.Println("Usage:\n     tz  t<zip>")
			return
		}
		zipName := args[0]
		list, err := tz.List(zipName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, s := range list {
			fmt.Println(s)
		}
		fmt.Printf("Total %d files\n", len(list))
	},
}
