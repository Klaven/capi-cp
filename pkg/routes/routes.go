package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	baseRoute = "/api/v1/"
)

func StartRouter() {
	router := mux.NewRouter()

	AddKubernetesRoutes(router)
	AddCapiRoutes(router)

	http.Handle("/", router)
}
