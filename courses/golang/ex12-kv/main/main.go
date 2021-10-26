package main

import (
	"context"
	"kv"
	"kv/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	storage := storage.NewStorage()
	r := mux.NewRouter()

	r.HandleFunc("/get", kv.GetHandler(storage)).Methods(http.MethodGet)
	r.HandleFunc("/add", kv.AddHandler(storage)).Methods(http.MethodPost)
	r.HandleFunc("/update", kv.UpdateHandler(storage)).Methods(http.MethodPost)
	r.HandleFunc("/delete", kv.DeleteHandler(storage)).Methods(http.MethodPost)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Println("Server started, hit Ctrl+C to stop")
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error: ", err)
	}
}
