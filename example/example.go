package example

//go:generate moq -out goimports -out mockexample_test.go . Example

type Args struct {
	a int
	b string
	c []byte
}

type Returns struct {
	e int
	f string
	g []byte
}

type Example interface {
	Helle(Args) Returns
}
