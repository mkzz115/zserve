package zserve

import (
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
    "github.com/gorilla/mux"
    "github.com/julienschmidt/httprouter"
    "github.com/mkzz115/zserve/common/log"
    "net/http"
    "time"
)

func NewZServer(name, addr string) *ZServe {
    return &ZServe{
        //ProcessorType: PROCESSOR_HTTP,
        Addr: addr,
        Name: name,
    }
}

type ZServe struct {
    Name          string
    ProcessorType string
    Addr          string
    //
    processor Processor
    fn        func(Configer) error
    cfg       Configer
}

func (t *ZServe) Init(cfger Configer, fn func(Configer) error, p Processor) error {
    log.Debug("ZServe.Init called")
    t.processor = p
    t.fn = fn
    t.cfg = cfger
    return nil
}

func (t *ZServe) Start() error {
    log.Debug("ZServe.Init called")
    // init depends serve
    if t.cfg != nil {
        //
        err := t.fn(t.cfg)
        if err != nil {
            return err
        }
    }
    err := t.processor.Init()
    if err != nil {
        return err
    }
    name, driver := t.processor.Driver()
    t.ProcessorType = name
    switch d := driver.(type) {
    case *httprouter.Router:
        return t.serveHttp(t.Addr, d)
    case *mux.Route:
        return t.serveHttp(t.Addr, d.GetHandler())
    case thrift.TProcessor:
        return t.serveThrift(t.Addr, d)
    default:
        return fmt.Errorf("processor:%s invalid", name)
    }
}

func (t *ZServe) serveThrift(addr string, processor thrift.TProcessor) error {
    log.Debug("ThriftServer.Init called")
    //handler := NewHelloServe()
    //process := hello.NewHelloServiceProcessor(handler)

    transport := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocol := thrift.NewTBinaryProtocolFactoryDefault()
    sock, err := thrift.NewTServerSocket(t.Addr)
    if err != nil {
        log.Error("ThriftServer.Start create socket error[%s], addr[%s]",
            err.Error(), t.Addr)
        return err
    }
    server := thrift.NewTSimpleServer4(processor, sock, transport, protocol)
    err = server.Listen()
    if err != nil {
        log.Error("ThriftServer.Start listen error[%s], addr[%s]",
            err.Error(), t.Addr)
        return err
    }
    return server.Serve()
}

func (t *ZServe) serveHttp(addr string, handle http.Handler) error {
    serv := &http.Server{
        Addr:              t.Addr,
        Handler:           handle,
        ReadHeaderTimeout: 5 * time.Second,
    }
    return serv.ListenAndServe()
}
