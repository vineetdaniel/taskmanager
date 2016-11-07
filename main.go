package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/vineetdaniel/taskmanager/common"
	"github.com/vineetdaniel/taskmanager/routers"
)

func main() {
	//common.Startup()

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    "127.0.0.1:8088",
		Handler: n,
	}
	log.Println("Listening....", common.AppConfig.Server)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
