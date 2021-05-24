package bootstrap

import (
	"cyc/goblog/pkg/route"
	"cyc/goblog/routes"
	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
