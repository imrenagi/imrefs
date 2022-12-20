package cli

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/imrenagi/imrefs"
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

			name := args[0]
			fileName := fmt.Sprintf(imrefs.FileFormat, name)
			data := args[1]
			conn, err := net.Dial("unix", fmt.Sprintf(imrefs.SockFormat, name))
			if err != nil {
				return err
			}
			defer conn.Close()

			r := imrefs.Request{
				Command:       imrefs.IPC_SEND,
				ContentLength: uint64(len(data)),
				Content:       io.NopCloser(bytes.NewBuffer([]byte(data))),
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
			fmt.Println(fmt.Sprintf("Data successfully written at %s", fileName))
			return nil
		},
	}
	return command
}
