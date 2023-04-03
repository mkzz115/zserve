package zserve

const (
	PROCESSOR_THRIFT = "thrift"
	PROCESSOR_HTTP   = "http"
)

type Serve interface {

	Init(Configer, func(Configer) error, Processor) error
	Start() error
}
