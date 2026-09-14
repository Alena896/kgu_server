package main

import (
	"fmt"
	"kgu_server/internal/handlers"
	"net/http"
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
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", handlers.ListExpensesHandler)
	http.HandleFunc("/add", handlers.AddExpenseHandler)
	http.ListenAndServe(":8080", nil)
}
