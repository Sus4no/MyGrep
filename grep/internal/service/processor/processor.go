package processor

import (
	"errors"
	"fmt"
	"grep/internal/adapter"
	"grep/internal/model/options"
	"grep/internal/service"
	"io"
)

type Processor struct {
	reader  adapter.Reader
	writer  adapter.Writer
	matcher adapter.Matcher
	options options.Options
}

func New(
	reader adapter.Reader,
	writer adapter.Writer,
	matcher adapter.Matcher,
	options options.Options,
) service.Processor {
	return &Processor{
		reader:  reader,
		writer:  writer,
		matcher: matcher,
		options: options,
	}
}

func (p *Processor) Run() error {
	for i := 1; true; i++ {
		line, err := p.reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			} else {
				return err
			}
		}

		match := p.matcher.MatchLine(line)
		if match != p.options.InverseResult {
			var err error
			if p.options.PrintStringNumber {
				err = p.writer.WriteLine(fmt.Sprintf("%d: %s", i, line))
			} else {
				err = p.writer.WriteLine(line)
			}
			if err != nil {
				return fmt.Errorf("writeLine: %w", err)
			}
		}

		p.writer.Flush()
	}

	return nil
}
