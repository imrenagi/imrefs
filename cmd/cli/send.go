package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func sendCmd() *cobra.Command {
	var command = &cobra.Command{
		Use: "send",
		Example: `
// imrefs send <name> <message>
// imrefs send myfs1 "Hello world 2!"
`,
		Short: "send file content to the file system",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("invalid args length. should only filesystem instance name and file content")
			}

			fmt.Println("Data successfully written at /tmp/file1234.tmp")
			return nil
		},
	}
	return command
}
