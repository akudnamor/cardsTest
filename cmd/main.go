package main

import (
	"BOARD/internal/http-server/handlers"
	"BOARD/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	storage, err := sqlite.New("storage/storage.db")

	router := chi.NewRouter()

	router.Handle("/internal/*", http.StripPrefix("/internal/", http.FileServer(http.Dir("internal"))))

	router.HandleFunc("/", handlers.AllCheck(storage))

	router.HandleFunc("/add", handlers.AddRandom(storage))

	router.HandleFunc("/order/{ID}", handlers.CheckOrder(storage))

	srv := http.Server{
		Addr:         "localhost:8081",
		Handler:      router,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("failed to start ", err)
	}
	log.Println("fail occured")
}
