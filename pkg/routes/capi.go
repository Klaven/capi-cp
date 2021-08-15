package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCapiRoutes(router *mux.Router) {
	router.HandleFunc(baseRoute+"my/first/capi/route", handleCapiRoute)
}

func handleCapiRoute(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "Something for sure")
}
