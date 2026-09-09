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

func createFastHTTPRequest(method, url string, useIOSUserAgent, useDefaultContentType bool, compressionMode string) *fasthttp.Request {
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod(method)
	request.SetRequestURI(url)

	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cache-Control", "no-store")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("X-Bloks-Version-Id", "02aa82d2510c3c91e0f953bebbdd0b36aedd15b580d2f6557e7b4b3b116ee5a3")
	request.Header.Set("User-Agent", "Instagram 410.1.0.63.71 Android (25/7.0; 120dpi; 720x540; OnePlus; HUAWEI-ZVPU; zb8688; vpudt3; en_US)")

	if useIOSUserAgent {
		request.Header.Set("User-Agent", "Instagram 361.0.0.46.88 (iPhone14,5; iOS 16_0_2; en_GB; en; scale=3.00; 1170x2532; 521442846) AppleWebKit/420+")
	}
	if useDefaultContentType {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if compressionMode != "" {
		request.Header.Set("Accept-Encoding", compressionMode)
	}
	return request
}

func createFastHTTPResponse(skipBody bool) *fasthttp.Response {
	response := fasthttp.AcquireResponse()
	response.SkipBody = skipBody
	return response
}
