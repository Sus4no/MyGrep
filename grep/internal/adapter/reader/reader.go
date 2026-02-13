package reader

import (
	"bufio"
	"fmt"
	"grep/internal/adapter"
	"io"
	"strings"
)

type Reader struct {
	reader *bufio.Reader
}

func New(in io.Reader) adapter.Reader {
	return &Reader{
		reader: bufio.NewReader(in),
	}
}

func (r *Reader) ReadLine() (string, error) {
	if r.reader == nil {
		return "", adapter.ErrNoSource
	}

	line, err := r.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("reading: %w", err)
	}

	return strings.TrimRight(line, "\r\n"), err
}
