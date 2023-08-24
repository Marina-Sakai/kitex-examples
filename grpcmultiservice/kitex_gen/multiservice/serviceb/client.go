// Code generated by Kitex v0.7.0. DO NOT EDIT.

package serviceb

import (
	"context"
	multiservice "github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ChatB(ctx context.Context, callOptions ...callopt.Option) (stream ServiceB_ChatBClient, err error)
}

type ServiceB_ChatBClient interface {
	streaming.Stream
	Send(*multiservice.Request) error
	CloseAndRecv() (*multiservice.Reply, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kServiceBClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kServiceBClient struct {
	*kClient
}

func (p *kServiceBClient) ChatB(ctx context.Context, callOptions ...callopt.Option) (stream ServiceB_ChatBClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatB(ctx)
}