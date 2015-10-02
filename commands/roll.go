package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ariejan/roll/core"
	"github.com/spf13/cobra"
	"github.com/spf13/hugo/utils"
)

// rollCmd is
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "roll a d20",
	Long: `roll is a command line utility to roll some dice

	roll is a fast and flexible dice roller, for you know, when you
	need to roll dice.

	Complete documentation is available at https://ariejan.net/roll`,
	Run: func(cmd *cobra.Command, args []string) {
		value := core.RollDie(20)
		fmt.Printf("You rolled 1d20: %d\n", value)
	},
}

func Execute() {
	AddCommands()
	utils.StopOnErr(rollCmd.Execute())
}

func AddCommands() {
	rollCmd.AddCommand(version)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
