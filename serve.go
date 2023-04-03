package zserve

import "zserve/common/confutil"

const (
	PROCESSOR_THRIFT = "thrift"
	PROCESSOR_HTTP   = "http"
)

type Serve interface {

	Init(*confutil.Config, func(DependsServer) error, Processor) error
	Start() error
}
