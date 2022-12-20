package imrefs

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

// IPC Command - 8 bit
// ContentLength - 64 bit
// Content

func ReadRequest(buf *bufio.Reader) (*Request, error) {
	r := &Request{}

	cb, err := buf.ReadByte()
	if err != nil {
		return nil, err
	}
	r.Command = cb

	err = binary.Read(buf, binary.BigEndian, &r.ContentLength)
	if err != nil {
		return nil, err
	}

	var nTotal int
	contentBuf := make([]byte, 1024)
	var content string
	var n int
	for {
		if uint64(nTotal) >= r.ContentLength {
			break
		}
		n, err = buf.Read(contentBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		content += string(contentBuf[:n])
		nTotal += n
	}

	bb := bytes.NewBufferString(content)
	r.Content = io.NopCloser(bb)
	return r, nil
}

type Command uint8

const (
	IPC_SEND uint8 = iota
	IPC_STOP
)

type Request struct {
	Command       uint8
	ContentLength uint64
	Content       io.ReadCloser
}

func (r *Request) Write(w io.Writer) error {
	b := []byte{r.Command}
	b = binary.BigEndian.AppendUint64(b, r.ContentLength)
	if _, err := w.Write(b); err != nil {
		return err
	}
	if r.ContentLength > 0 {
		if _, err := io.Copy(w, r.Content); err != nil {
			return err
		}
	}

	return nil
}
