package handlers

import (
	"cardsTest/internal/storage/sqlite"
	"cardsTest/lib/random"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

func AllCheck(st *sqlite.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		pr, err := st.GetAllProduct()
		if err != nil {
			fmt.Errorf("failed to get prod", err)
		}

		t, err := template.ParseFiles("internal/templates/index.html")
		if err != nil {
			fmt.Errorf("failed to parse template", err)
		}

		err = t.Execute(w, pr)
		if err != nil {
			log.Println(err)
		}

	}
}

func AddFiveRandom(st *sqlite.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		for i := 0; i < 5; i++ {
			err := st.AddProduct(random.RandomInt(), random.RandomString(), random.RandomFloat64())
			if err != nil {
				fmt.Errorf("failed to add prod", err)
			}
		}

		w.Write([]byte("ADDED 5"))

	}
}
