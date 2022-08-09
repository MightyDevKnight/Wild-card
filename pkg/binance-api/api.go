package bapi

import (
	"fmt"
	"github.com/1makarov/binance-nft-buy/internal/domain/account"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Api struct {
	request *fasthttp.Request
	http    *fasthttp.Client
}

type Api_Demo struct {
	request *fasthttp.Request
	http    *fasthttp.Client
}

func New(setting account.Setting) (*Api, error) {
	http := initHttp(setting.Proxy)
	request, err := initHeaders(setting.BAuth)
	if err != nil {
		return nil, err
	}
	return &Api{http: http, request: request}, nil
}

func initHttp(proxy string) *fasthttp.Client {
	c := &fasthttp.Client{}
	if proxy != "" {
		c.Dial = fasthttpproxy.FasthttpHTTPDialer(proxy)
		return c
	}

	request, err := initHeaders(setting.BAuth)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func initHeaders(bAuth *account.BAuth) (*fasthttp.Request, error) {
	if bAuth.Cookie == "" || bAuth.Csrf == "" {
		return nil, fmt.Errorf("empty field .env")
	}
	return r, nil
}
