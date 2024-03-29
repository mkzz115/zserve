// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package order

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mkzz115/zserve/common/pub/idl/gen-go/base"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = base.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - UID
type GetWishListReq struct {
	UID int64 `thrift:"Uid,1" json:"Uid"`
}

func NewGetWishListReq() *GetWishListReq {
	return &GetWishListReq{}
}

func (p *GetWishListReq) GetUID() int64 {
	return p.UID
}
func (p *GetWishListReq) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetWishListReq) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.UID = v
	}
	return nil
}

func (p *GetWishListReq) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetWishListReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetWishListReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Uid", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:Uid: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.UID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.Uid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:Uid: ", p), err)
	}
	return err
}

func (p *GetWishListReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetWishListReq(%+v)", *p)
}

type WishListData struct {
}

func NewWishListData() *WishListData {
	return &WishListData{}
}

func (p *WishListData) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WishListData) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("WishListData"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WishListData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WishListData(%+v)", *p)
}

// Attributes:
//  - Items
//  - ErrInfo
type GetWishListRes struct {
	Items []*WishListData `thrift:"items,1" json:"items,omitempty"`
	// unused fields # 2 to 99
	ErrInfo *base.ErrorInfo `thrift:"ErrInfo,100" json:"ErrInfo,omitempty"`
}

func NewGetWishListRes() *GetWishListRes {
	return &GetWishListRes{}
}

var GetWishListRes_Items_DEFAULT []*WishListData

func (p *GetWishListRes) GetItems() []*WishListData {
	return p.Items
}

var GetWishListRes_ErrInfo_DEFAULT *base.ErrorInfo

func (p *GetWishListRes) GetErrInfo() *base.ErrorInfo {
	if !p.IsSetErrInfo() {
		return GetWishListRes_ErrInfo_DEFAULT
	}
	return p.ErrInfo
}
func (p *GetWishListRes) IsSetItems() bool {
	return p.Items != nil
}

func (p *GetWishListRes) IsSetErrInfo() bool {
	return p.ErrInfo != nil
}

func (p *GetWishListRes) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 100:
			if err := p.readField100(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetWishListRes) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*WishListData, 0, size)
	p.Items = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &WishListData{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Items = append(p.Items, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetWishListRes) readField100(iprot thrift.TProtocol) error {
	p.ErrInfo = &base.ErrorInfo{}
	if err := p.ErrInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ErrInfo), err)
	}
	return nil
}

func (p *GetWishListRes) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetWishListRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField100(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetWishListRes) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetItems() {
		if err := oprot.WriteFieldBegin("items", thrift.LIST, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:items: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Items)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Items {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:items: ", p), err)
		}
	}
	return err
}

func (p *GetWishListRes) writeField100(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrInfo() {
		if err := oprot.WriteFieldBegin("ErrInfo", thrift.STRUCT, 100); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ErrInfo: ", p), err)
		}
		if err := p.ErrInfo.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ErrInfo), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ErrInfo: ", p), err)
		}
	}
	return err
}

func (p *GetWishListRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetWishListRes(%+v)", *p)
}
