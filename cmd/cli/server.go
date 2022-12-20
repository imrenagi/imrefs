package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func serverCmd() *cobra.Command {
	var command = &cobra.Command{
		Use: "server",
		Example: `
// imrefs server
`,
		Short: "start file system server",
		RunE: func(c *cobra.Command, args []string) error {

			f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
			if err != nil {
				return err
			}
			defer f.Close()

			timer := time.NewTimer(20 * time.Second)

		caw:
			for {
				select {
				case t := <-time.After(1 * time.Second):
					f.WriteString(fmt.Sprintf("%v\n", t.String()))
				case <-timer.C:
					break caw
				}
			}
			return nil
		},
	}
	return command
}
