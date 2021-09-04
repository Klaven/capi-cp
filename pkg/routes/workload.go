package routes

import (
	"github.com/gorilla/mux"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	workloadRoute = "/workload"
)

func getWorkloadBaseRoute() string {
	return baseRouteV1 + workloadRoute
}

type WorkloadRouter struct {
	Client client.Client
}

func NewWorkloadRouter(client client.Client) *WorkloadRouter {
	return &WorkloadRouter{Client: client}
}

func (c *WorkloadRouter) AddRoutes(router *mux.Router) {}
