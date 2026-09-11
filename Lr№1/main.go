package main

import (
	"fmt"
	"net/http"
)

func homeHandler(conclusion http.ResponseWriter, input *http.Request) {
	if input.Method != "GET" {
		http.Error(conclusion, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Fprintf(conclusion, "Привет! Твой сервер на Go работает!")
	}
}

func aboutHandler(conclusion http.ResponseWriter, input *http.Request) {
	if input.Method != "GET" {
		http.Error(conclusion, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Fprintf(conclusion, "Проект: Основы Go и HTTP-сервера")
	}
}

func pingHandler(conclusion http.ResponseWriter, input *http.Request) {
	if input.Method != "GET" {
		http.Error(conclusion, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Fprintf(conclusion, "200 ОК pong")
	}
}

func main() {
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}
