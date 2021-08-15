package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AddKubernetesRoutes(router *mux.Router) {
	router.HandleFunc(baseRoute+"my/first/kubernetes/route", handleKubeRoute)
}

func handleKubeRoute(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "Something for sure")
}
