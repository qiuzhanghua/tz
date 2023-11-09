package main

import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "tz",
		Short: "TZ",
		Long:  "Use Zip to compress/extract/list files",
	}
	rootCmd.AddCommand(TzXCmd, TzCCmd, TzTCmd)
	cobra.CheckErr(rootCmd.Execute())

}
