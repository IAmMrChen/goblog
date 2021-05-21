package route

import (
	"github.com/gorilla/mux"
)

func Name2URL(routeName string, pair ...string) string  {
	var Router *mux.Router
	url, err := Router.Get(routeName).URL(pair...)

	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}