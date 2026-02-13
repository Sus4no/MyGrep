package matcher

import (
	"grep/internal/adapter"
	"regexp"
)

type Matcher struct {
	pattern *regexp.Regexp
}

func New(pattern string, caseSensetive bool) (adapter.Matcher, error) {
	if !caseSensetive {
		pattern = "(?i)" + pattern
	}
	exp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, adapter.ErrInvalidPattern
	}
	return &Matcher{
		pattern: exp,
	}, nil
}

func (m *Matcher) MatchLine(line string) bool {
	return m.pattern.MatchString(line)
}
