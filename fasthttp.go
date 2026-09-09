package main

import (
	"crypto/tls"
	"net"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

func createProxyClient(connections int) (*fasthttp.Client, chan string) {
	channel := make(chan string)
	return &fasthttp.Client{
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS13},
		Dial: func(addr string) (net.Conn, error) {
			return fasthttpproxy.FasthttpHTTPDialer(<-channel)(addr)
		},
		MaxConnsPerHost: connections,
	}, channel
}

func createRequest(method, url, encoding string) *fasthttp.Request {
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(method)
	request.SetRequestURI(url)
	request.Header.Set("User-Agent", "GraphQL monitor example")
	if encoding != "" {
		request.Header.Set("Accept-Encoding", encoding)
	}
	return request
}
