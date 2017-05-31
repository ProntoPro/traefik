package middlewares

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

// Detect proxy protocol
type ProxyProtocol struct {
	Handler negroni.Handler
}

// ServerHTTP is a function used by negroni
func NewProxyProtocolParser() *ProxyProtocol {
	proxyProtocol := ProxyProtocol{}

	proxyProtocol.Handler = negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		r.Header["X-Forwarded-For"] = []string{"213.23.21.23"}

		next.ServeHTTP(w, r)
	})

	return &proxyProtocol
}
