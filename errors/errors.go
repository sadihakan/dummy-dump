package errors

import "github.com/sadihakan/dummy-dump/model"

type DummyDumpError struct {
	Status int
	Errors interface{}
	Detail string
	s      string
}

func New(err model.DummyDumpError, args ...interface{}) error {
	e := &DummyDumpError{s: string(err)}

	if len(args) >= 1 {
		e.Status = args[0].(int)
	}

	if len(args) >= 2 {
		e.Detail = args[1].(string)
	}

	if len(args) >= 3 {
		e.Errors = args[2].(interface{})
	}

	return e
}

func (e *DummyDumpError) Error() string {
	return e.s
}