package main

import (
	"cardsTest/internal/http-server/handlers"
	"cardsTest/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	storage, err := sqlite.New("storage/storage.db")

	router := chi.NewRouter()

	router.HandleFunc("/", handlers.AllCheck(storage))

	router.HandleFunc("/add", handlers.AddFiveRandom(storage))

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("internal"))))

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
