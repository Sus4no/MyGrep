package cli

import (
	"grep/internal/model/options"
)

type Params struct {
	Options  options.Options
	Filepath string
	Pattern  string
}
