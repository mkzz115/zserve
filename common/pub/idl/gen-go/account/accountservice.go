// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package account

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

type AccountService interface {
	// Parameters:
	//  - Req
	UserLogIn(req *LogInReq) (r *LogInRes, err error)
	// Parameters:
	//  - Req
	UserLogOut(req *LogOutReq) (r *LogOutRes, err error)
	// Parameters:
	//  - Req
	UserRegister(req *RegisterReq) (r *RegisterRes, err error)
	// Parameters:
	//  - Req
	UserChangePw(req *ChangePwReq) (r *ChangePwRes, err error)
}

type AccountServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAccountServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AccountServiceClient {
	return &AccountServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAccountServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AccountServiceClient {
	return &AccountServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Req
func (p *AccountServiceClient) UserLogIn(req *LogInReq) (r *LogInRes, err error) {
	if err = p.sendUserLogIn(req); err != nil {
		return
	}
	return p.recvUserLogIn()
}

func (p *AccountServiceClient) sendUserLogIn(req *LogInReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("UserLogIn", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AccountServiceUserLogInArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AccountServiceClient) recvUserLogIn() (value *LogInRes, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "UserLogIn" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "UserLogIn failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "UserLogIn failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "UserLogIn failed: invalid message type")
		return
	}
	result := AccountServiceUserLogInResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Req
func (p *AccountServiceClient) UserLogOut(req *LogOutReq) (r *LogOutRes, err error) {
	if err = p.sendUserLogOut(req); err != nil {
		return
	}
	return p.recvUserLogOut()
}

func (p *AccountServiceClient) sendUserLogOut(req *LogOutReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("UserLogOut", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AccountServiceUserLogOutArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AccountServiceClient) recvUserLogOut() (value *LogOutRes, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "UserLogOut" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "UserLogOut failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "UserLogOut failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "UserLogOut failed: invalid message type")
		return
	}
	result := AccountServiceUserLogOutResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Req
func (p *AccountServiceClient) UserRegister(req *RegisterReq) (r *RegisterRes, err error) {
	if err = p.sendUserRegister(req); err != nil {
		return
	}
	return p.recvUserRegister()
}

func (p *AccountServiceClient) sendUserRegister(req *RegisterReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("UserRegister", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AccountServiceUserRegisterArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AccountServiceClient) recvUserRegister() (value *RegisterRes, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "UserRegister" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "UserRegister failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "UserRegister failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "UserRegister failed: invalid message type")
		return
	}
	result := AccountServiceUserRegisterResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Req
func (p *AccountServiceClient) UserChangePw(req *ChangePwReq) (r *ChangePwRes, err error) {
	if err = p.sendUserChangePw(req); err != nil {
		return
	}
	return p.recvUserChangePw()
}

func (p *AccountServiceClient) sendUserChangePw(req *ChangePwReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("UserChangePw", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AccountServiceUserChangePwArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AccountServiceClient) recvUserChangePw() (value *ChangePwRes, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "UserChangePw" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "UserChangePw failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "UserChangePw failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "UserChangePw failed: invalid message type")
		return
	}
	result := AccountServiceUserChangePwResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type AccountServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AccountService
}

func (p *AccountServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AccountServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AccountServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAccountServiceProcessor(handler AccountService) *AccountServiceProcessor {

	self8 := &AccountServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["UserLogIn"] = &accountServiceProcessorUserLogIn{handler: handler}
	self8.processorMap["UserLogOut"] = &accountServiceProcessorUserLogOut{handler: handler}
	self8.processorMap["UserRegister"] = &accountServiceProcessorUserRegister{handler: handler}
	self8.processorMap["UserChangePw"] = &accountServiceProcessorUserChangePw{handler: handler}
	return self8
}

func (p *AccountServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type accountServiceProcessorUserLogIn struct {
	handler AccountService
}

func (p *accountServiceProcessorUserLogIn) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountServiceUserLogInArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("UserLogIn", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AccountServiceUserLogInResult{}
	var retval *LogInRes
	var err2 error
	if retval, err2 = p.handler.UserLogIn(args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing UserLogIn: "+err2.Error())
		oprot.WriteMessageBegin("UserLogIn", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("UserLogIn", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type accountServiceProcessorUserLogOut struct {
	handler AccountService
}

func (p *accountServiceProcessorUserLogOut) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountServiceUserLogOutArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("UserLogOut", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AccountServiceUserLogOutResult{}
	var retval *LogOutRes
	var err2 error
	if retval, err2 = p.handler.UserLogOut(args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing UserLogOut: "+err2.Error())
		oprot.WriteMessageBegin("UserLogOut", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("UserLogOut", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type accountServiceProcessorUserRegister struct {
	handler AccountService
}

func (p *accountServiceProcessorUserRegister) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountServiceUserRegisterArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("UserRegister", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AccountServiceUserRegisterResult{}
	var retval *RegisterRes
	var err2 error
	if retval, err2 = p.handler.UserRegister(args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing UserRegister: "+err2.Error())
		oprot.WriteMessageBegin("UserRegister", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("UserRegister", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type accountServiceProcessorUserChangePw struct {
	handler AccountService
}

func (p *accountServiceProcessorUserChangePw) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AccountServiceUserChangePwArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("UserChangePw", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AccountServiceUserChangePwResult{}
	var retval *ChangePwRes
	var err2 error
	if retval, err2 = p.handler.UserChangePw(args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing UserChangePw: "+err2.Error())
		oprot.WriteMessageBegin("UserChangePw", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("UserChangePw", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Req
type AccountServiceUserLogInArgs struct {
	Req *LogInReq `thrift:"req,1" json:"req"`
}

func NewAccountServiceUserLogInArgs() *AccountServiceUserLogInArgs {
	return &AccountServiceUserLogInArgs{}
}

var AccountServiceUserLogInArgs_Req_DEFAULT *LogInReq

func (p *AccountServiceUserLogInArgs) GetReq() *LogInReq {
	if !p.IsSetReq() {
		return AccountServiceUserLogInArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *AccountServiceUserLogInArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AccountServiceUserLogInArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AccountServiceUserLogInArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = &LogInReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *AccountServiceUserLogInArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserLogIn_args"); err != nil {
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

func (p *AccountServiceUserLogInArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *AccountServiceUserLogInArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserLogInArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AccountServiceUserLogInResult struct {
	Success *LogInRes `thrift:"success,0" json:"success,omitempty"`
}

func NewAccountServiceUserLogInResult() *AccountServiceUserLogInResult {
	return &AccountServiceUserLogInResult{}
}

var AccountServiceUserLogInResult_Success_DEFAULT *LogInRes

func (p *AccountServiceUserLogInResult) GetSuccess() *LogInRes {
	if !p.IsSetSuccess() {
		return AccountServiceUserLogInResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AccountServiceUserLogInResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountServiceUserLogInResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *AccountServiceUserLogInResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &LogInRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountServiceUserLogInResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserLogIn_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *AccountServiceUserLogInResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountServiceUserLogInResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserLogInResult(%+v)", *p)
}

// Attributes:
//  - Req
type AccountServiceUserLogOutArgs struct {
	Req *LogOutReq `thrift:"req,1" json:"req"`
}

func NewAccountServiceUserLogOutArgs() *AccountServiceUserLogOutArgs {
	return &AccountServiceUserLogOutArgs{}
}

var AccountServiceUserLogOutArgs_Req_DEFAULT *LogOutReq

func (p *AccountServiceUserLogOutArgs) GetReq() *LogOutReq {
	if !p.IsSetReq() {
		return AccountServiceUserLogOutArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *AccountServiceUserLogOutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AccountServiceUserLogOutArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AccountServiceUserLogOutArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = &LogOutReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *AccountServiceUserLogOutArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserLogOut_args"); err != nil {
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

func (p *AccountServiceUserLogOutArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *AccountServiceUserLogOutArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserLogOutArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AccountServiceUserLogOutResult struct {
	Success *LogOutRes `thrift:"success,0" json:"success,omitempty"`
}

func NewAccountServiceUserLogOutResult() *AccountServiceUserLogOutResult {
	return &AccountServiceUserLogOutResult{}
}

var AccountServiceUserLogOutResult_Success_DEFAULT *LogOutRes

func (p *AccountServiceUserLogOutResult) GetSuccess() *LogOutRes {
	if !p.IsSetSuccess() {
		return AccountServiceUserLogOutResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AccountServiceUserLogOutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountServiceUserLogOutResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *AccountServiceUserLogOutResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &LogOutRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountServiceUserLogOutResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserLogOut_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *AccountServiceUserLogOutResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountServiceUserLogOutResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserLogOutResult(%+v)", *p)
}

// Attributes:
//  - Req
type AccountServiceUserRegisterArgs struct {
	Req *RegisterReq `thrift:"req,1" json:"req"`
}

func NewAccountServiceUserRegisterArgs() *AccountServiceUserRegisterArgs {
	return &AccountServiceUserRegisterArgs{}
}

var AccountServiceUserRegisterArgs_Req_DEFAULT *RegisterReq

func (p *AccountServiceUserRegisterArgs) GetReq() *RegisterReq {
	if !p.IsSetReq() {
		return AccountServiceUserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *AccountServiceUserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AccountServiceUserRegisterArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AccountServiceUserRegisterArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = &RegisterReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *AccountServiceUserRegisterArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserRegister_args"); err != nil {
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

func (p *AccountServiceUserRegisterArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *AccountServiceUserRegisterArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserRegisterArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AccountServiceUserRegisterResult struct {
	Success *RegisterRes `thrift:"success,0" json:"success,omitempty"`
}

func NewAccountServiceUserRegisterResult() *AccountServiceUserRegisterResult {
	return &AccountServiceUserRegisterResult{}
}

var AccountServiceUserRegisterResult_Success_DEFAULT *RegisterRes

func (p *AccountServiceUserRegisterResult) GetSuccess() *RegisterRes {
	if !p.IsSetSuccess() {
		return AccountServiceUserRegisterResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AccountServiceUserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountServiceUserRegisterResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *AccountServiceUserRegisterResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &RegisterRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountServiceUserRegisterResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserRegister_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *AccountServiceUserRegisterResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountServiceUserRegisterResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserRegisterResult(%+v)", *p)
}

// Attributes:
//  - Req
type AccountServiceUserChangePwArgs struct {
	Req *ChangePwReq `thrift:"req,1" json:"req"`
}

func NewAccountServiceUserChangePwArgs() *AccountServiceUserChangePwArgs {
	return &AccountServiceUserChangePwArgs{}
}

var AccountServiceUserChangePwArgs_Req_DEFAULT *ChangePwReq

func (p *AccountServiceUserChangePwArgs) GetReq() *ChangePwReq {
	if !p.IsSetReq() {
		return AccountServiceUserChangePwArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *AccountServiceUserChangePwArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AccountServiceUserChangePwArgs) Read(iprot thrift.TProtocol) error {
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

func (p *AccountServiceUserChangePwArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = &ChangePwReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *AccountServiceUserChangePwArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserChangePw_args"); err != nil {
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

func (p *AccountServiceUserChangePwArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *AccountServiceUserChangePwArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserChangePwArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AccountServiceUserChangePwResult struct {
	Success *ChangePwRes `thrift:"success,0" json:"success,omitempty"`
}

func NewAccountServiceUserChangePwResult() *AccountServiceUserChangePwResult {
	return &AccountServiceUserChangePwResult{}
}

var AccountServiceUserChangePwResult_Success_DEFAULT *ChangePwRes

func (p *AccountServiceUserChangePwResult) GetSuccess() *ChangePwRes {
	if !p.IsSetSuccess() {
		return AccountServiceUserChangePwResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AccountServiceUserChangePwResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountServiceUserChangePwResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *AccountServiceUserChangePwResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &ChangePwRes{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountServiceUserChangePwResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UserChangePw_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *AccountServiceUserChangePwResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountServiceUserChangePwResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountServiceUserChangePwResult(%+v)", *p)
}
