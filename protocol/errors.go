package protocol

import "errors"

var (
	//ErrNilResult is of type error
	ErrNilResult = errors.New("result is nil")
	//ErrUnknown is of type string which returns unknown error
	ErrUnknown = ProxyError{"Unknown Error,instance is not selected, error is nil"}
	//ErrUnExpectedHandlerChainResponse is of type string which returns unexpected handler error
	ErrUnExpectedHandlerChainResponse = ProxyError{"Response from Handler chain is nil,better to check if handler chain is empty, or some handler just return a nil response"}
)

//ProxyError is a struct
type ProxyError struct {
	Message string
}

func (e ProxyError) Error() string {
	return e.Message
}
