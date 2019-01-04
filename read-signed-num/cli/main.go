package main

import (
	"time"

	"code.byted.org/gopkg/thrift"
	"code.byted.org/inf/ByteGraph/bgdb/log"
	"github.com/lemonwx/some_test/read-signed-num/idl/echo"
)

func main() {
	socket, _ := thrift.NewTSocketTimeout(":1234", time.Second)
	transport := thrift.NewTFramedTransport(socket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport.Open()
	cli := echo.NewEchoServiceClientFactory(transport, protocolFactory)
	resp, err := cli.Echo(&echo.EchoReq{Id: -1, TypeA1: -2})
	log.Debug(resp, err)
}
