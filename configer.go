package zserve

import "github.com/mkzz115/zserve/common/confutil"

type Configer interface {
    GetConfig() *confutil.Config

}

type DepConfiger interface {
    GetConfig() Configer

}