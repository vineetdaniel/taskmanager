package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/vineetdaniel/AiOps/apiv1/common"
	"github.com/vineetdaniel/AiOps/apiv1/routers"
)

func main() {
	//common.Startup()

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening....", common.AppConfig.Server)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
