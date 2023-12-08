package server

import (
	"context"
)

type MockFIlter struct {
}

func (c MockFIlter) Handle(ctx context.Context, req, rsp interface{},
	next func(ctx context.Context, reqbody interface{}, rspbody interface{}) error) error {

	return nil
}

func (c MockFIlter) Filter(ctx context.Context, req interface{},
	next interface{}) (
	rspbody interface{}, err error) {

	return rspbody, nil
}

type FilterFunc func(reqBody interface{}) (MockFIlter, error)

// ServiceDesc describes a proto service.
type ServiceDesc struct {
	ServiceName  string
	HandlerType  interface{}
	Methods      interface{}
	Streams      interface{}
	StreamHandle interface{}
}

type Method struct {
	Name     string
	Func     interface{}
	Bindings interface{}
}

type Service struct {
}

func (s Service) Register(interface{}, interface{}) error {
	return nil
}
