package doproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func DoProxy(p *httputil.ReverseProxy, target *url.URL) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host
		p.ServeHTTP(w, r)
	}
}
