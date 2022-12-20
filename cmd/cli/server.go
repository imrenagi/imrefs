package cli

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/imrenagi/imrefs"
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
			fileName := fmt.Sprintf("files/file-%s.tmp", name)
			l, err := net.Listen("unix", fmt.Sprintf("files/file-%s.sock", name))
			if err != nil {
				return err
			}

			for {
				conn, err := l.Accept()
				if err != nil {
					return err
				}

				go handleConn(conn, fileName)
			}
			return nil
		},
	}
	return command
}

func handleConn(conn net.Conn, fileName string) {
	req, err := imrefs.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		conn.Write([]byte("FAILED"))
		return
	}

	switch req.Command {
	case imrefs.IPC_SEND:
		writeToFile(conn, fileName, req.Content)
	case imrefs.IPC_STOP:
		// TODO properly stop server
		conn.Write([]byte("OK"))
		os.Exit(0)
	}
}

func writeToFile(conn net.Conn, name string, content io.ReadCloser) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, content)
	if err != nil {
		return nil
	}
	conn.Write([]byte("OK"))
	return nil
}
