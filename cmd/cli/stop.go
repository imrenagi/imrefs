package cli

import (
	"fmt"
	"net"

	"github.com/imrenagi/imrefs"
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
			name := args[0]
			conn, err := net.Dial("unix", fmt.Sprintf(imrefs.SockFormat, name))
			if err != nil {
				return err
			}
			defer conn.Close()

			r := imrefs.Request{
				Command: imrefs.IPC_STOP,
			}

			err = r.Write(conn)
			if err != nil {
				return err
			}

			buf := make([]byte, 1024)
			_, err = conn.Read(buf)
			if err != nil {
				return err
			}

			fmt.Printf("Filesystem %s stopped\n", name)
			return nil
		},
	}
	return command
}
