package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	baseRoute = "/api/v1"
)

func StartRouter(client client.Client) *mux.Router {
	router := mux.NewRouter()

	AddKubernetesRoutes(router)

	capiRouter := NewCapiRouter(client)

	capiRouter.AddCapiRoutes(router)

	http.Handle("/", router)

	return router
}
