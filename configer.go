package zserve

import "zserve/common/confutil"

type Configer interface {
    GetConfig() *confutil.Config

}

type DepConfiger interface {
    GetConfig() Configer

}