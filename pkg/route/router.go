package route

import "github.com/gorilla/mux"

var Router *mux.Router

func Initialize()  {
	Router = mux.NewRouter()
}

func Name2URL(routeName string, pair ...string) string  {
	url, err := Router.Get(routeName).URL(pair...)

	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}