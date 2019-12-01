// +build !go1.11

package middleware

import (
	"net/http"
	"net/http/httputil"

	"github.com/matm/echo/v4"
)

func proxyHTTP(t *ProxyTarget, c echo.Context, config ProxyConfig) http.Handler {
	return httputil.NewSingleHostReverseProxy(t.URL)
}
