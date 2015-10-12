package commands

import (
	"fmt"

	"github.com/ariejan/roll/core"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Roll",
	Long:  `This is my version, there are many like it, but this one is mine.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(core.Version)
	},
}
