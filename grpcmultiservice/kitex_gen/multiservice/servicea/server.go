// Code generated by Kitex v0.6.2. DO NOT EDIT.
package servicea

import (
	multiservice "github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler multiservice.ServiceA, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
