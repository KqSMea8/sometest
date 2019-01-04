package main

import (
	"code.byted.org/gopkg/thrift"
	"github.com/lemonwx/log"
	"github.com/lemonwx/some_test/read-signed-num/idl/echo"
)

type EchoService struct {
}

func (e *EchoService) Echo(req *echo.EchoReq) (r *echo.EchoResp, err error) {
	log.Debug(req.Id, req.TypeA1)
	return &echo.EchoResp{}, nil
}

func main() {
	echoSvrTransport, err := thrift.NewTServerSocket(":1234")
	if err != nil {
		log.Fatalf("master admin service transport establish at %s failed: %v", ":1234", err)
	}
	echoSvr := &EchoService{}
	echoSvrProcessor := echo.NewEchoServiceProcessor(echoSvr)
	echoSvrTransportFactory := thrift.NewTFramedTransportFactory(thrift.NewTBufferedTransportFactory(8192))
	echoSvrProtoFactory := thrift.NewTBinaryProtocolFactoryDefault()
	echoSvrSvr := thrift.NewTSimpleServer4(echoSvrProcessor, echoSvrTransport,
		echoSvrTransportFactory, echoSvrProtoFactory)
	echoSvrSvr.Serve()
}
