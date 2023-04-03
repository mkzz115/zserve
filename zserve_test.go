package zserve

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"testing"
	"zserve/common/confutil"
	"zserve/common/log"
	"zserve/common/pub/adapter"
	"zserve/common/pub/idl/gen-go/hello"
)

func TestZServe_Start(t *testing.T) {
	log.Init("./test.log")
	go test_thrift(t)
	test_http(t)
}


func test_thrift(t *testing.T) {
	fn := func(cf Configer) error {
		t.Log("fnnnnnnn")
		return nil
	}
	proc := &thriftServe{}
	zserve := NewZServer("test_thrift", "127.0.0.1:12306")

	zserve.Init(nil, fn, proc)
	zserve.Start()
}

type thriftServe struct{}

func (p *thriftServe) Init() error {
	println("thriftServe.Init")
	return nil
}


func NewHelloServe() *HelloServe {
	return &HelloServe{}
}

var dps DependsServer

func InitHandle(dp DependsServer) error {
	_ = log.Init("InitHandle called")
	dps = dp
	return nil
}

type HelloServe struct {
	dp DependsServer
}

func (h *HelloServe) GetHello(req *hello.HelloReq) (r *hello.HelloRes, err error) {
	log.Info("HelloServe.GetHello called, req.uid[%d]", req.GetUID())
	data := req.GetUID()
	res := hello.NewHelloRes()
	res.Ack = &data
	var uid int64 = 0
	//err = dps.DbInstance().Get(&uid, "select uid from user_account_t where account=?", "zz")
	//if err != nil {
	//	log.Info("select db err:%s", err.Error())
	//	return res, err
	//}
	//log.Info("account:%d", uid)
	res.Ack = &uid
	return res, nil
}


func (p *thriftServe) Driver() (string, interface{}) {
	handler := NewHelloServe()
	processor := hello.NewHelloServiceProcessor(handler)
	return PROCESSOR_THRIFT, processor
}

func test_http(t *testing.T) {
	fn := func(cf Configer) error {
		t.Log("fnnnnnnn")
		return nil
	}
	proc := &httpServe{}
	zserve := NewZServer("test_http", "127.0.0.1:12307")

	zserve.Init(nil, fn, proc)
	zserve.Start()

}

type httpServe struct{}

func (p *httpServe) Init() error {
	println("httpServe.Init")
	return nil
}



func NewLogicHandle(cfg *confutil.Config) *logicInfo {
	return &logicInfo{
		cfg: cfg,
	}
}

type logicInfo struct {
	cfg *confutil.Config
}

func (m *logicInfo) Hello(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("hello is called, data[%v]", string(buf))
	req := &hello.HelloReq{}
	err = json.Unmarshal(buf, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := adapter.GetHelloAdapter(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resData)
}


func (p *httpServe) Driver() (string, interface{}) {
	r := httprouter.New()

	h := NewLogicHandle(nil)
	r.HandlerFunc(http.MethodPost, "/hello", h.Hello)
	return PROCESSOR_HTTP, r
}
