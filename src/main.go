package main

import (
	"log"
	"net/http"
	"tech-challenge/registry"
	"tech-challenge/router"
	"time"
)

const (
	PORT = "8080"
)

func main() {
	portEnv := PORT

	//datastore := repository.NewKVStoreLocal()
	registry := registry.NewRegistry()
	router := router.Start(registry.NewAppController())

	srv := &http.Server{
		Addr:         ":" + portEnv,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen:%+s\n", err)
	}
}
