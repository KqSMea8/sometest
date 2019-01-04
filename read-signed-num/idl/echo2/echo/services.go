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

type EchoServiceEchoArgs struct {
	Req *EchoReq `thrift:"req,1" json:"req"`
}

func NewEchoServiceEchoArgs() *EchoServiceEchoArgs {
	return &EchoServiceEchoArgs{
		Req: NewEchoReq(),
	}
}

func (p *EchoServiceEchoArgs) writeFields(w *thrift.BufferWriter) error {
	// Write p.Req
	{
		w.WriteFieldHeader(thrift.TType_STRUCT, 1)
		if err := p.Req.Write(w); err != nil {
			return err
		}
	}
	w.WriteInt8(thrift.TType_STOP)
	return nil
}
func (p *EchoServiceEchoArgs) Write(w *thrift.BufferWriter) error {
	return p.writeFields(w)
}

func (p *EchoServiceEchoArgs) readFields(r *thrift.BufferReader) error {
	for {
		t, err := r.ReadInt8()
		if err != nil {
			return err
		}
		if t == thrift.TType_STOP {
			break
		}
		fieldId, err := r.ReadInt16()
		if err != nil {
			return err
		}
		switch fieldId {
		case 1: // Read p.Req
			if t != thrift.TType_STRUCT {
				return thrift.ErrFieldType
			}
			p.Req = NewEchoReq()
			if err := p.Req.Read(r); err != nil {
				return err
			}
		default:
			if err := r.SkipField(t); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *EchoServiceEchoArgs) Read(r *thrift.BufferReader) error {
	return p.readFields(r)
}

func (p *EchoServiceEchoArgs) WriteMessage(w *thrift.BufferWriter, seq int32) (err error) {
	w.WriteMessageHeader("Echo", thrift.CALL, seq)
	return p.Write(w)
}

func (p *EchoServiceEchoReturns) WriteMessage(w *thrift.BufferWriter, seq int32) (err error) {
	w.WriteMessageHeader("Echo", thrift.REPLY, seq)
	return p.Write(w)
}

type EchoServiceEchoReturns struct {
	Return *EchoResp `thrift:"Return,0" json:"Return"`
}

func NewEchoServiceEchoReturns() *EchoServiceEchoReturns {
	return &EchoServiceEchoReturns{
		Return: NewEchoResp(),
	}
}

func (p *EchoServiceEchoReturns) writeFields(w *thrift.BufferWriter) error {
	// Write p.Return
	{
		w.WriteFieldHeader(thrift.TType_STRUCT, 0)
		if err := p.Return.Write(w); err != nil {
			return err
		}
	}
	w.WriteInt8(thrift.TType_STOP)
	return nil
}
func (p *EchoServiceEchoReturns) Write(w *thrift.BufferWriter) error {
	return p.writeFields(w)
}

func (p *EchoServiceEchoReturns) readFields(r *thrift.BufferReader) error {
	for {
		t, err := r.ReadInt8()
		if err != nil {
			return err
		}
		if t == thrift.TType_STOP {
			break
		}
		fieldId, err := r.ReadInt16()
		if err != nil {
			return err
		}
		switch fieldId {
		case 0: // Read p.Return
			if t != thrift.TType_STRUCT {
				return thrift.ErrFieldType
			}
			p.Return = NewEchoResp()
			if err := p.Return.Read(r); err != nil {
				return err
			}
		default:
			if err := r.SkipField(t); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *EchoServiceEchoReturns) Read(r *thrift.BufferReader) error {
	return p.readFields(r)
}

type EchoServiceClient struct {
	thrift.Client
}

func NewEchoServiceClient(cli thrift.Client) *EchoServiceClient {
	return &EchoServiceClient{Client: cli}
}

func (_cli *EchoServiceClient) Echo(ctx context.Context, req *EchoReq) (*EchoResp, error) {
	EchoArgs := EchoServiceEchoArgs{
		Req: req,
	}
	EchoReturns := NewEchoServiceEchoReturns()
	_err := _cli.Client.Invoke(ctx, "Echo", &EchoArgs, EchoReturns)
	if _err != nil {
		return nil, _err
	}
	return EchoReturns.Return, nil
}

type EchoServiceHandler interface {
	Echo(ctx context.Context, req *EchoReq) (*EchoResp, error)
}

type EchoServiceServerHandler struct {
	h EchoServiceHandler
}

func NewEchoServiceServerHandler(h EchoServiceHandler) EchoServiceServerHandler {
	return EchoServiceServerHandler{h: h}
}

func (h EchoServiceServerHandler) Serve(ctx context.Context, r *thrift.BufferReader, w *thrift.BufferWriter) (err error) {
	type request struct {
		method string
		seq    int32
		args   interface{}
	}
	ch := make(chan request, MaxServerPipeline)
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	defer func() { close(ch); cancel(); wg.Wait() }()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for req := range ch {
			switch req.method {
			case "Echo":
				{
					args := req.args.(*EchoServiceEchoArgs)
					ret, err := h.h.Echo(ctx, args.Req)
					returns := NewEchoServiceEchoReturns()
					if ret != nil {
						returns.Return = ret
					}
					if err != nil {
						thrift.WriteApplicationException(w, err)
						continue
					}
					if err := returns.WriteMessage(w, req.seq); err != nil {
						return
					}
				} //  Echo
			}
			if err := w.Flush(); err != nil {
				return
			}
		}
	}()
	for {
		r.Clear()
		var req request
		name, tid, seq, err := r.ReadMessageHeader()
		if err != nil {
			return err
		}
		if tid != thrift.CALL && tid != thrift.ONEWAY {
			return thrift.ErrMessageType
		}
		req.method = name
		req.seq = seq
		switch name {
		case "Echo":
			{
				args := NewEchoServiceEchoArgs()
				if err := args.Read(r); err != nil {
					return err
				}
				req.args = args
			} //  Echo
		default:
			return thrift.ErrUnknownFunction
		}
		ch <- req
	}
}
