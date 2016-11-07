package routers

import "github.com/gorilla/mux"
import "github.com/vineetdaniel/AiOps/apiv1/controllers"

func SetUrlRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/urls", controllers.GetUrls).Methods("GET")
	router.HandleFunc("/urls/create", controllers.CreateUrl).Methods("POST")
	return router
}
