package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/imrenagi/imrefs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func serverCmd() *cobra.Command {
	var command = &cobra.Command{
		Use: "server",
		Example: `
// imrefs server myfs1
`,
		Short: "start file system server",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("invalid args length. should only accept one argument")
			}
			name := args[0]

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)

			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			signal.Notify(ch, syscall.SIGTERM)

			go func() {
				oscall := <-ch
				log.Warn().Msgf("system call:%+v", oscall)
				cancel()
			}()

			f := imrefs.New(name)
			if err := f.Run(ctx); err != nil {
				return err
			}

			return nil
		},
	}
	return command
}
