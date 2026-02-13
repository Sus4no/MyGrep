package adapter

type Reader interface {
	ReadLine() (string, error)
}
