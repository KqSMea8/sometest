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

type EchoReq struct {
	Id   int64 `thrift:"id,1" json:"id"`
	Type int32 `thrift:"type,2" json:"type"`
}

func NewEchoReq() *EchoReq {
	return &EchoReq{}
}

func (p *EchoReq) writeFields(w *thrift.BufferWriter) error {
	// Write p.Id
	if p.Id != 0 {
		w.WriteFieldHeader(thrift.TType_I64, 1)
		w.WriteInt64(p.Id)
	}
	// Write p.Type
	if p.Type != 0 {
		w.WriteFieldHeader(thrift.TType_I32, 2)
		w.WriteInt32(p.Type)
	}
	w.WriteInt8(thrift.TType_STOP)
	return nil
}
func (p *EchoReq) Write(w *thrift.BufferWriter) error {
	return p.writeFields(w)
}

func (p *EchoReq) readFields(r *thrift.BufferReader) error {
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
		case 1: // Read p.Id
			if t != thrift.TType_I64 {
				return thrift.ErrFieldType
			}
			v0, err := r.ReadInt64()
			if err != nil {
				return err
			}
			p.Id = v0
		case 2: // Read p.Type
			if t != thrift.TType_I32 {
				return thrift.ErrFieldType
			}
			v0, err := r.ReadInt32()
			if err != nil {
				return err
			}
			p.Type = v0
		default:
			if err := r.SkipField(t); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *EchoReq) Read(r *thrift.BufferReader) error {
	return p.readFields(r)
}

type EchoResp struct {
	Id   int64 `thrift:"id,1" json:"id"`
	Type int32 `thrift:"type,2" json:"type"`
}

func NewEchoResp() *EchoResp {
	return &EchoResp{}
}

func (p *EchoResp) writeFields(w *thrift.BufferWriter) error {
	// Write p.Id
	if p.Id != 0 {
		w.WriteFieldHeader(thrift.TType_I64, 1)
		w.WriteInt64(p.Id)
	}
	// Write p.Type
	if p.Type != 0 {
		w.WriteFieldHeader(thrift.TType_I32, 2)
		w.WriteInt32(p.Type)
	}
	w.WriteInt8(thrift.TType_STOP)
	return nil
}
func (p *EchoResp) Write(w *thrift.BufferWriter) error {
	return p.writeFields(w)
}

func (p *EchoResp) readFields(r *thrift.BufferReader) error {
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
		case 1: // Read p.Id
			if t != thrift.TType_I64 {
				return thrift.ErrFieldType
			}
			v0, err := r.ReadInt64()
			if err != nil {
				return err
			}
			p.Id = v0
		case 2: // Read p.Type
			if t != thrift.TType_I32 {
				return thrift.ErrFieldType
			}
			v0, err := r.ReadInt32()
			if err != nil {
				return err
			}
			p.Type = v0
		default:
			if err := r.SkipField(t); err != nil {
				return err
			}
		}
	}
	return nil
}
func (p *EchoResp) Read(r *thrift.BufferReader) error {
	return p.readFields(r)
}
