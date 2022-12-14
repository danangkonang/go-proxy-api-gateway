package app

import (
	"fmt"
	"net/http"

	"github.com/danangkonang/go-proxy-api-gateway/helper"
	"github.com/danangkonang/go-proxy-api-gateway/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 404, "page not found", nil)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 405, "Method NotAllowed", nil)
	})

	router.ProductRouter(r)

	header := []string{"*"}
	method := []string{"*"}
	origin := []string{"*"}
	fmt.Println("local server started at http://localhost:7000")
	http.ListenAndServe(":7000", handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))
}
