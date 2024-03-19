package handlers

import (
	"BOARD/internal/storage/sqlite"
	"BOARD/lib/random"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func CheckOrder(st *sqlite.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(chi.URLParam(r, "ID"))
		if err != nil {
			fmt.Errorf("failed to conv string", err)
		}

		pr, err := st.GetProductByID(ID)
		if err != nil {
			fmt.Errorf("failed to get prod", err)
		}

		w.Write([]byte(pr.Name))
		return
	}
}

func AddRandom(st *sqlite.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := st.AddProduct(random.RandomInt(), random.RandomString(), random.RandomFloat64())
		if err != nil {
			fmt.Errorf("failed to add prod", err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
