package cli

import (
	"errors"
	"flag"
	"fmt"
	"grep/internal/adapter"
	"grep/internal/adapter/matcher"
	readerAdapter "grep/internal/adapter/reader"
	"grep/internal/adapter/writer"
	"grep/internal/lib/ptr"
	model "grep/internal/model/options"
	"grep/internal/service/processor"
	"os"
)

var (
	ErrTooManyArgs = errors.New("too many args")
	ErrTooFewArgs  = errors.New("too few args")
	ErrInvalidFile = errors.New("invalid file")
)

func Run() {
	params, err := parseArgs()
	if err != nil {
		ExitWithError(err.Error())
	}

	var reader adapter.Reader
	if params.Filepath != "" {
		file, err := os.Open(params.Filepath)
		if err != nil {
			ExitWithError(ErrInvalidFile.Error())
		}
		defer file.Close()
		reader = readerAdapter.New(file)
	} else {
		reader = readerAdapter.New(os.Stdin)
	}

	writer := writer.New(os.Stdout)

	matcher, err := matcher.New(params.Pattern, params.Options.IgnoreRegister)
	if err != nil {
		ExitWithError(err.Error())
	}

	proc := processor.New(reader, writer, matcher, params.Options)

	err = proc.Run()
	if err != nil {
		ExitWithError(err.Error())
	}
}

func parseArgs() (Params, error) {
	n := flag.Bool("n", false, "Print string number")
	i := flag.Bool("i", false, "Ignore register")
	v := flag.Bool("v", false, "Inverse result (show non-matching)")
	flag.Parse()

	options := model.Options{
		PrintStringNumber: ptr.ZeroIfNil(n),
		IgnoreRegister:    ptr.ZeroIfNil(i),
		InverseResult:     ptr.ZeroIfNil(v),
	}

	args := flag.Args()

	var filepath, pattern string
	switch {
	case len(args) > 2:
		return Params{}, ErrTooManyArgs
	case len(args) == 2:
		pattern = args[0]
		filepath = args[1]
	case len(args) == 1:
		pattern = args[0]
	default:
		return Params{}, ErrTooFewArgs
	}

	return Params{
		Options:  options,
		Filepath: filepath,
		Pattern:  pattern,
	}, nil
}

func ExitWithError(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s", msg)
	os.Exit(1)
}
