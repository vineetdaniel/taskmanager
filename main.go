package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/vineetdaniel/taskmanager/common"
)

func main() {
	common.Startup()

	router := router.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Sever,
		Handler: n,
	}
	log.Println("Listening....")
	server.ListenAndServe()
}
