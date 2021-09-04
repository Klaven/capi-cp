package routes

import (
	"github.com/gorilla/mux"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	providersRoute = "/providers"
)

func getProvidersBaseRoute() string {
	return baseRouteV1 + providersRoute
}

type ProvidersRouter struct {
	Client client.Client
}

func NewProvidersRouter(client client.Client) *ProvidersRouter {
	return &ProvidersRouter{Client: client}
}

func (c *ProvidersRouter) AddRoutes(router *mux.Router) {}
