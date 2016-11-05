package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/vineetdaniel/taskmanager/common"
	"github.com/vineetdaniel/taskmanager/routers"
)

func main() {
	common.Startup()

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Sever,
		Handler: n,
	}
	log.Println("Listening....")
	server.ListenAndServe()
}
