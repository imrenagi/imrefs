package cli

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var command = &cobra.Command{
		Use:   "imrefs",
		Short: "file system for programming challenge",
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}
	command.AddCommand(initCmd(), sendCmd(), stopCmd(), serverCmd())
	return command
}
