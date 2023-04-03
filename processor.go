package zserve

type Processor interface {
	Init() error
	Driver() (string, interface{})
}
