package router

import (
	"net/http/httputil"
	"net/url"

	"github.com/danangkonang/go-proxy-api-gateway/doproxy"
	"github.com/gorilla/mux"
)

func ProductRouter(r *mux.Router) {
	target, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(target)
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/", doproxy.DoProxy(proxy, target)).Methods("GET")
}
