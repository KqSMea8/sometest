// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// GENERATOR VERSION RPC/THRIFT b61d3c2
package echo

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"code.byted.org/rpc/thrift/pkg/thrift"
)

var (
	_ = context.Canceled
	_ = errors.New
	_ = fmt.Sprintf
	_ = sync.WaitGroup{}

	_ = thrift.ThriftPackageIsVersion102
)

var (
	MaxMapElements    int32 = 1 << 20
	MaxSetElements    int32 = 1 << 20
	MaxListElements   int32 = 1 << 20
	MaxServerPipeline int32 = 10
)

var GoUnusedProtection__ struct{}
