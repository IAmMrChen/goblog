package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func SetRoute(r *mux.Router)  {
	Router = r
}

func Name2URL(routeName string, pair ...string) string  {
	fmt.Println(pair)
	url, err := Router.Get(routeName).URL(pair...)

	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}