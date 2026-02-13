package adapter

type Matcher interface {
	MatchLine(string) bool
}
