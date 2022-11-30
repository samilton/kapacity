package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var ()

func init() {
	rootCmd.AddCommand(describeCmd)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a specific resource",
	Long:  "Provide detail on a specific resource",

	Run: func(cmd *cobra.Command, args []string) {
		describe()
	},
}

func describe() {
	fmt.Println("Describing a thing")
}
