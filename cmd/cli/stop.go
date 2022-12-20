package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func stopCmd() *cobra.Command {
	var command = &cobra.Command{
		Use: "stop",
		Example: `
// imrefs stop <name>
// imrefs stop myfs1
`,
		Short: "stop process",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("invalid args length. should only accept one argument")
			}

			fmt.Println("Filesystem myfs1 stopped")
			return nil
		},
	}
	return command
}
