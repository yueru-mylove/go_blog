package main

import (
	"awesomeProject/package/settings"
	"awesomeProject/route"
	"fmt"
	"net/http"
)


func main() {

	router := route.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}
