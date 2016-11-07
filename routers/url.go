package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vineetdaniel/AiOps/apiv1/common"
	"github.com/vineetdaniel/AiOps/apiv1/controllers"
)

func SetUrlRoutes(router *mux.Router) *mux.Router {
	urlRouter := mux.NewRouter()
	urlRouter.HandleFunc("/urls", controllers.GetUrls).Methods("GET")
	urlRouter.HandleFunc("/urls/create", controllers.CreateUrl).Methods("POST")
	router.PathPrefix("/urls").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(urlRouter),
	))
	return router
}
