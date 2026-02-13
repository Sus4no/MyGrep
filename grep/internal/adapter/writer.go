package adapter

type Writer interface {
	WriteLine(string) error
	Flush() error
}
