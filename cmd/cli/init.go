package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	var command = &cobra.Command{
		Use: "init",
		Example: `
// imrefs init <name>
// imrefs init myfs1
`,
		Short: "initialize file system",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("invalid args length. should only accept one argument")
			}

			lsCmd := exec.Command("go", "run", "cmd/main.go", "server")
			err := lsCmd.Start()
			if err != nil {
				return err
			}

			pid := lsCmd.Process.Pid
			name := args[0]
			fileName := fmt.Sprintf("files/file-%d-%s.tmp", pid, name)

			f, err := os.OpenFile(fileName, os.O_CREATE, 0777)
			if err != nil {
				return err
			}
			defer f.Close()

			fmt.Printf("Filesystem %s successfully created at %s with PID %d\n", name, fileName, pid)
			return nil
		},
	}
	return command
}
