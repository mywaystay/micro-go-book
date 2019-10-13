package config

import (
	"github.com/go-kit/kit/log"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/bootstrap"
	_ "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/bootstrap"
	conf "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/config"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	_ "github.com/openzipkin/zipkin-go/reporter/recorder"
	"github.com/spf13/viper"
	"os"
	"sync"
)

const (
	kConfigType = "CONFIG_TYPE"
)

var ZipkinTracer *zipkin.Tracer
var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()

	if err := conf.LoadRemoteConfig(); err != nil {
		Logger.Log("Fail to load remote config", err)
	}

	if err := conf.Sub("mysql", &conf.MysqlConfig); err != nil {
		Logger.Log("Fail to parse mysql", err)
	}
	if err := conf.Sub("trace", &conf.TraceConfig); err != nil {
		Logger.Log("Fail to parse trace", err)
	}
	zipkinUrl := "http://" + conf.TraceConfig.Host + ":" + conf.TraceConfig.Port + conf.TraceConfig.Url
	Logger.Log("zipkin url", zipkinUrl)
	initTracer(zipkinUrl)
}

func initDefault() {
	viper.SetDefault(kConfigType, "yaml")
}

func initTracer(zipkinURL string) {
	var (
		err           error
		useNoopTracer = zipkinURL == ""
		reporter      = zipkinhttp.NewReporter(zipkinURL)
	)
	//defer reporter.Close()
	zEP, _ := zipkin.NewEndpoint(bootstrap.HttpConfig.ServiceName, bootstrap.HttpConfig.Port)
	ZipkinTracer, err = zipkin.NewTracer(
		reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
	)
	if err != nil {
		Logger.Log("err", err)
		os.Exit(1)
	}
	if !useNoopTracer {
		Logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinURL)
	}
}

type SecResult struct {
	ProductId int    `json:"product_id"` //商品ID
	UserId    int    `json:"user_id"`    //用户ID
	Token     string `json:"token"`      //Token
	TokenTime int64  `json:"token_time"` //Token生成时间
	Code      int    `json:"code"`       //状态码
}

type SecRequest struct {
	ProductId     int             `json:"product_id"` //商品ID
	Source        string          `json:"source"`
	AuthCode      string          `json:"auth_code"`
	SecTime       string          `json:"sec_time"`
	Nance         string          `json:"nance"`
	UserId        int             `json:"user_id"`
	UserAuthSign  string          `json:"user_auth_sign"` //用户授权签名
	AccessTime    int64           `json:"access_time"`
	ClientAddr    string          `json:"client_addr"`
	ClientRefence string          `json:"client_refence"`
	CloseNotify   <-chan bool     `json:"-"`
	ResultChan    chan *SecResult `json:"-"`
}

var SkAppContext = &SkAppCtx{
	UserConnMap: make(map[string]chan *SecResult, 1024),
	SecReqChan:  make(chan *SecRequest, 1024),
}

type SkAppCtx struct {
	SecReqChan       chan *SecRequest
	SecReqChanSize   int
	RWSecProductLock sync.RWMutex

	UserConnMap     map[string]chan *SecResult
	UserConnMapLock sync.Mutex
}

const (
	ProductStatusNormal       = 0 //商品状态正常
	ProductStatusSaleOut      = 1 //商品售罄
	ProductStatusForceSaleOut = 2 //商品强制售罄
)
