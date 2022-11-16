package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wacli",
		Short: "wacli is a tool for winter aop",
		Long:  `wacli is a tool for winter aop`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("please run: wacli -h")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
