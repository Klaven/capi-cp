package main

import (
	"flag"
	"net/http"

	"github.com/rancher/wrangler/pkg/kubeconfig"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/klaven/capi-cp/pkg/routes"
)

var (
	kubeconfigfile string
	Scheme         = runtime.NewScheme()
)

func init() {
	flag.StringVar(&kubeconfigfile, "c", "", "Path to a kubeconfig")
	flag.Parse()
}

func main() {
	client, err := getKubeClient(kubeconfigfile, Scheme)
	if err != nil {
		return
	}
	router := routes.StartRouter(client)
	http.ListenAndServe(":8080", router)
}

func getKubeClient(kconfig string, scheme *runtime.Scheme) (client.Client, error) {
	cfg, err := kubeconfig.GetNonInteractiveClientConfig(kconfig).ClientConfig()
	if err != nil {
		return nil, err
	}
	return client.New(cfg, client.Options{Scheme: scheme})
}
