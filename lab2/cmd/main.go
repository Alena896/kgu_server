package main

import (
	"database/sql"
	"fmt"
	"kgu_server/internal/handlers"
	"kgu_server/internal/repository"
	"net/http"

	"os"

	_ "modernc.org/sqlite"
)

func aboutHandler(conclusion http.ResponseWriter, input *http.Request) {
	if input.Method != "GET" {
		http.Error(conclusion, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Fprintf(conclusion, "Проект: Основы Go и HTTP-сервера")
	}
}

func main() {
	db, err := sql.Open("sqlite", "expenses.db")
	if err != nil {
		panic(err)
	}
	queryBytes, err := os.ReadFile("table")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(queryBytes))
	if err != nil {
		panic(err)
	}
	repo := repository.NewExpenseRepository(db)
	h := &handlers.ExpenseHandler{Repo: repo}
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", h.ListExpensesHandler)
	http.HandleFunc("/add", h.AddExpenseHandler)
	http.HandleFunc("/delete", h.DeleteExpenseHandler)
	http.HandleFunc("/edit", h.EditExpenseHandler)
	http.ListenAndServe(":8080", nil)
}
