package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/types"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	capiRoute = "/capi"
)

func getCapiBaseRoute() string {
	return baseRoute + capiRoute
}

type CapiRouter struct {
	Client client.Client
}

func NewCapiRouter(client client.Client) *CapiRouter {
	return &CapiRouter{Client: client}
}

func (c *CapiRouter) AddCapiRoutes(router *mux.Router) {
	router.HandleFunc(getCapiBaseRoute()+"/clusters", c.handleGetClusters).Methods("GET")
}

func handleCapiRoute(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "Something for sure")
}

func (c *CapiRouter) handleGetClusters(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	clusters := &v1alpha4.ClusterList{}
	c.Client.List(context.Background(), clusters)
	for _, v := range clusters.Items {
		json.NewEncoder(responseWriter).Encode(v) // TODO: return a sane amount of data about clusters
		return
	}
	responseWriter.WriteHeader(http.StatusNotFound)
	json.NewEncoder(responseWriter).Encode("No Cluster Created Yet!!!")
}

func (c *CapiRouter) handleGetCluster(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(request)
	clusterName := vars["name"]
	clusterNamespace := vars["namespace"]

	cluster := &v1alpha4.Cluster{}
	err := c.Client.Get(context.Background(), types.NamespacedName{Name: clusterName, Namespace: clusterNamespace}, cluster)
	if err == nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(responseWriter).Encode("Cluster Not Found")
	}

	json.NewEncoder(responseWriter).Encode(c)
}
