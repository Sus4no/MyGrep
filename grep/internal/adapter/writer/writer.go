package writer

import (
	"bufio"
	"fmt"
	"grep/internal/adapter"
	"io"
)

type Writer struct {
	writer *bufio.Writer
}

func New(out io.Writer) adapter.Writer {
	return &Writer{
		writer: bufio.NewWriter(out),
	}
}

func (w *Writer) WriteLine(line string) error {
	if _, err := w.writer.WriteString(line); err != nil {
		return fmt.Errorf("writing string: %w", err)
	}
	if err := w.writer.WriteByte('\n'); err != nil {
		return fmt.Errorf("writing byte: %w", err)
	}
	return nil
}

func (w *Writer) Flush() error {
	if err := w.writer.Flush(); err != nil {
		return fmt.Errorf("flushing: %w", err)
	}
	return nil
}
