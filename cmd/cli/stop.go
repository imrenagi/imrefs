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
			conn, err := net.Dial("unix", fmt.Sprintf("files/file-%s.sock", name))
			if err != nil {
				return err
			}
			defer conn.Close()

			r := imrefs.Request{
				Command:       imrefs.IPC_STOP,
				ContentLength: 0,
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

			fmt.Println("Filesystem myfs1 stopped")
			return nil
		},
	}
	return command
}
