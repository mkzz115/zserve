package zserve

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"zserve/common/confutil"
	"zserve/common/idgenutil"
)

type DependsServer interface {
	GenId() (int64, error)
	//DbInstance() *sqlx.DB
}

func NewDepServe(dfger DepConfiger) DependsServer {
	ds := &dependsDerve{
		cfg: dfger.GetConfig().GetConfig(),
	}
	ds.idgen = idgenutil.NewIdWorker()
	return ds
}

type dependsDerve struct {
	cfg   *confutil.Config
	idgen idgenutil.IdGenerator
}

func (d *dependsDerve) GenId() (int64, error) {
	return d.idgen.GenId()
}

func (d *dependsDerve) DbInstance() *sqlx.DB {
	if len(d.cfg.Depends.MysqlConfig.Dsn) == 0 {
		return nil
	}
	db, err := sqlx.Open("mysql", d.cfg.Depends.MysqlConfig.Dsn)
	if err != nil {
		panic(err)
	}
	return db
}
