// Code generated by Kitex v0.6.2. DO NOT EDIT.

package serviceb

import (
	"context"
	"fmt"
	multiservice "github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return serviceBServiceInfo
}

var serviceBServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ServiceB"
	handlerType := (*multiservice.ServiceB)(nil)
	methods := map[string]kitex.MethodInfo{
		"ChatB": kitex.NewMethodInfo(chatBHandler, newChatBArgs, newChatBResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "multiservice",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func chatBHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &serviceBChatBServer{st}
	return handler.(multiservice.ServiceB).ChatB(stream)
}

type serviceBChatBClient struct {
	streaming.Stream
}

func (x *serviceBChatBClient) Send(m *multiservice.Request) error {
	return x.Stream.SendMsg(m)
}
func (x *serviceBChatBClient) CloseAndRecv() (*multiservice.Reply, error) {
	if err := x.Stream.Close(); err != nil {
		return nil, err
	}
	m := new(multiservice.Reply)
	return m, x.Stream.RecvMsg(m)
}

type serviceBChatBServer struct {
	streaming.Stream
}

func (x *serviceBChatBServer) SendAndClose(m *multiservice.Reply) error {
	return x.Stream.SendMsg(m)
}

func (x *serviceBChatBServer) Recv() (*multiservice.Request, error) {
	m := new(multiservice.Request)
	return m, x.Stream.RecvMsg(m)
}

func newChatBArgs() interface{} {
	return &ChatBArgs{}
}

func newChatBResult() interface{} {
	return &ChatBResult{}
}

type ChatBArgs struct {
	Req *multiservice.Request
}

func (p *ChatBArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(multiservice.Request)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatBArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatBArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatBArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChatBArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChatBArgs) Unmarshal(in []byte) error {
	msg := new(multiservice.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatBArgs_Req_DEFAULT *multiservice.Request

func (p *ChatBArgs) GetReq() *multiservice.Request {
	if !p.IsSetReq() {
		return ChatBArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatBArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChatBArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChatBResult struct {
	Success *multiservice.Reply
}

var ChatBResult_Success_DEFAULT *multiservice.Reply

func (p *ChatBResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(multiservice.Reply)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatBResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatBResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatBResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChatBResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChatBResult) Unmarshal(in []byte) error {
	msg := new(multiservice.Reply)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatBResult) GetSuccess() *multiservice.Reply {
	if !p.IsSetSuccess() {
		return ChatBResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatBResult) SetSuccess(x interface{}) {
	p.Success = x.(*multiservice.Reply)
}

func (p *ChatBResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChatBResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ChatB(ctx context.Context) (ServiceB_ChatBClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "ChatB", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &serviceBChatBClient{res.Stream}
	return stream, nil
}
