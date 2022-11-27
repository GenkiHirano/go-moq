package example

import "testing"

func TestExample(t *testing.T) {
	mock := &ExampleMock{
		HelleFunc: func(args Args) Returns{
			return Returns{}
		},
	}
	mock.Helle(Args{})
}
