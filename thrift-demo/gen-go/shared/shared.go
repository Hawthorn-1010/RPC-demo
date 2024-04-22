// Code generated by Thrift Compiler (0.20.0). DO NOT EDIT.

package shared

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = bytes.Equal
var _ = context.Background
var _ = errors.New
var _ = fmt.Printf
var _ = slog.Log
var _ = time.Now
var _ = thrift.ZERO
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

// Attributes:
//  - Key
//  - Value
type SharedStruct struct {
  Key int32 `thrift:"key,1" db:"key" json:"key"`
  Value string `thrift:"value,2" db:"value" json:"value"`
}

func NewSharedStruct() *SharedStruct {
  return &SharedStruct{}
}


func (p *SharedStruct) GetKey() int32 {
  return p.Key
}

func (p *SharedStruct) GetValue() string {
  return p.Value
}
func (p *SharedStruct) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedStruct)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *SharedStruct)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *SharedStruct) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "SharedStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedStruct) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedStruct) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "value", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Value)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  return err
}

func (p *SharedStruct) Equals(other *SharedStruct) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Key != other.Key { return false }
  if p.Value != other.Value { return false }
  return true
}

func (p *SharedStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedStruct(%+v)", *p)
}

func (p *SharedStruct) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*shared.SharedStruct",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*SharedStruct)(nil)

func (p *SharedStruct) Validate() error {
  return nil
}
type SharedService interface {
  // Parameters:
  //  - Key
  GetStruct(ctx context.Context, key int32) (_r *SharedStruct, _err error)
}

type SharedServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewSharedServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SharedServiceClient {
  return &SharedServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewSharedServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SharedServiceClient {
  return &SharedServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewSharedServiceClient(c thrift.TClient) *SharedServiceClient {
  return &SharedServiceClient{
    c: c,
  }
}

func (p *SharedServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *SharedServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *SharedServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Key
func (p *SharedServiceClient) GetStruct(ctx context.Context, key int32) (_r *SharedStruct, _err error) {
  var _args0 SharedServiceGetStructArgs
  _args0.Key = key
  var _result2 SharedServiceGetStructResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "getStruct", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  if _ret3 := _result2.GetSuccess(); _ret3 != nil {
    return _ret3, nil
  }
  return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "getStruct failed: unknown result")
}

type SharedServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler SharedService
}

func (p *SharedServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *SharedServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *SharedServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewSharedServiceProcessor(handler SharedService) *SharedServiceProcessor {

  self4 := &SharedServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["getStruct"] = &sharedServiceProcessorGetStruct{handler:handler}
return self4
}

func (p *SharedServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x5.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x5

}

type sharedServiceProcessorGetStruct struct {
  handler SharedService
}

func (p *sharedServiceProcessorGetStruct) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  var _write_err6 error
  args := SharedServiceGetStructArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "getStruct", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelCauseFunc
    ctx, cancel = context.WithCancelCause(ctx)
    defer cancel(nil)
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelCauseFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel(thrift.ErrAbandonRequest)
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := SharedServiceGetStructResult{}
  if retval, err2 := p.handler.GetStruct(ctx, args.Key); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
    if errors.Is(err2, thrift.ErrAbandonRequest) {
      return false, thrift.WrapTException(err2)
    }
    if errors.Is(err2, context.Canceled) {
      if err := context.Cause(ctx); errors.Is(err, thrift.ErrAbandonRequest) {
        return false, thrift.WrapTException(err)
      }
    }
    _exc7 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getStruct: " + err2.Error())
    if err2 := oprot.WriteMessageBegin(ctx, "getStruct", thrift.EXCEPTION, seqId); err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := _exc7.Write(ctx, oprot); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := oprot.WriteMessageEnd(ctx); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := oprot.Flush(ctx); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if _write_err6 != nil {
      return false, thrift.WrapTException(_write_err6)
    }
    return true, err
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 := oprot.WriteMessageBegin(ctx, "getStruct", thrift.REPLY, seqId); err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := result.Write(ctx, oprot); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := oprot.WriteMessageEnd(ctx); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := oprot.Flush(ctx); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if _write_err6 != nil {
    return false, thrift.WrapTException(_write_err6)
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
type SharedServiceGetStructArgs struct {
  Key int32 `thrift:"key,1" db:"key" json:"key"`
}

func NewSharedServiceGetStructArgs() *SharedServiceGetStructArgs {
  return &SharedServiceGetStructArgs{}
}


func (p *SharedServiceGetStructArgs) GetKey() int32 {
  return p.Key
}
func (p *SharedServiceGetStructArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedServiceGetStructArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *SharedServiceGetStructArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "getStruct_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "key", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *SharedServiceGetStructArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedServiceGetStructArgs(%+v)", *p)
}

func (p *SharedServiceGetStructArgs) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*shared.SharedServiceGetStructArgs",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*SharedServiceGetStructArgs)(nil)

// Attributes:
//  - Success
type SharedServiceGetStructResult struct {
  Success *SharedStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSharedServiceGetStructResult() *SharedServiceGetStructResult {
  return &SharedServiceGetStructResult{}
}

var SharedServiceGetStructResult_Success_DEFAULT *SharedStruct
func (p *SharedServiceGetStructResult) GetSuccess() *SharedStruct {
  if !p.IsSetSuccess() {
    return SharedServiceGetStructResult_Success_DEFAULT
  }
  return p.Success
}
func (p *SharedServiceGetStructResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *SharedServiceGetStructResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *SharedServiceGetStructResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &SharedStruct{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *SharedServiceGetStructResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "getStruct_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SharedServiceGetStructResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *SharedServiceGetStructResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SharedServiceGetStructResult(%+v)", *p)
}

func (p *SharedServiceGetStructResult) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*shared.SharedServiceGetStructResult",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*SharedServiceGetStructResult)(nil)


