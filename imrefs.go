package imrefs

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

var SockFormat = "files/file-%s.sock"
var FileFormat = "files/file-%s.tmp"

type FS struct {
	name string

	stopCh chan error
}

func New(name string) *FS {
	fs := &FS{
		name:   name,
		stopCh: make(chan error),
	}
	return fs
}

func (f *FS) fileName() string {
	return fmt.Sprintf(FileFormat, f.name)
}

func (f *FS) Run(ctx context.Context) error {
	l, err := net.Listen("unix", fmt.Sprintf(SockFormat, f.name))
	if err != nil {
		return err
	}
	defer l.Close()

	log.Info().Msg("listener created")

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				f.stopCh <- err
				return
			}
			go f.handleConn(conn)
		}
	}()

	log.Info().Msg("listener run")

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-f.stopCh:
		return err
	}
}

func (f *FS) handleConn(conn net.Conn) {
	log.Debug().Msg("handling connection")

	req, err := ReadRequest(bufio.NewReader(conn))
	if err != nil {
		conn.Write([]byte("FAILED"))
		return
	}

	log.Debug().Msg("get the request")

	switch req.Command {
	case IPC_SEND:
		f.write(conn, req.Content)
	case IPC_STOP:
		conn.Write([]byte("OK"))
		f.stop()
	}
}

func (f *FS) write(conn net.Conn, content io.ReadCloser) error {

	file, err := os.OpenFile(f.fileName(), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return nil
	}
	conn.Write([]byte("OK"))
	return nil
}

func (f *FS) stop() {
	f.stopCh <- nil
}
