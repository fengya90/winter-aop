package cmd

import (
	"fmt"
	"github.com/fengya90/winter-aop/wacli/gen"

	"github.com/spf13/cobra"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate AOP code",
		Long: `
Generate AOP code from configuration files or parameters.
For example:

	wacli gen -d /myhome/mydir1 -d /myhome/mydir2
	wacli gen -f /myhome/myconfig`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(sourceDirs) == 0 && configFilePath == "" {
				fmt.Println("source_dir or  config_file_path is required")
				return
			}
			gen.GenFunc(sourceDirs, configFilePath)
		},
	}
)

func init() {
	genCmd.Flags().StringArrayVarP(&sourceDirs, "source_dir", "d", []string{}, "the source code directories")
	genCmd.Flags().StringVarP(&configFilePath, "config_file_path", "f", "", "configuration file for the source code directories")
	rootCmd.AddCommand(genCmd)
}
