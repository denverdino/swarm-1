package api

import (
	"crypto/tls"
	"net/http"
)

func Hijack(tlsConfig *tls.Config, addr string, w http.ResponseWriter, r *http.Request) error {
	return hijack(tlsConfig, addr, w, r)
}

func ProxyAsync(tlsConfig *tls.Config, addr string, w http.ResponseWriter, r *http.Request, callback func(*http.Response)) error {
	return proxyAsync(tlsConfig, addr, w, r, callback)
}

func CloseIdleConnections(client *http.Client) {
	closeIdleConnections(client)
}
