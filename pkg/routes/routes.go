package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	baseRouteV1 = "/api/v1"
)

func StartRouter(client client.Client) *mux.Router {
	router := mux.NewRouter()

	AddKubernetesRoutes(router)

	mgmtRouter := NewMgmtRouter(client)
	mgmtRouter.AddRoutes(router)

	workloadRouter := NewWorkloadRouter(client)
	workloadRouter.AddRoutes(router)

	providersRouter := NewProvidersRouter(client)
	providersRouter.AddRoutes(router)

	http.Handle("/", router)

	return router
}
