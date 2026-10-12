package handlers

import (
	"fmt"
	"html/template"
	"kgu_server/internal/models"
	"kgu_server/internal/repository"
	"net/http"
	"strconv"
)

type ExpenseHandler struct {
	Repo *repository.ExpenseRepository
}

func (h *ExpenseHandler) ListExpensesHandler(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, l := template.ParseFiles("web/templates/layout.html", "web/templates/index.html")
	fmt.Println(l)
	tmpl.ExecuteTemplate(w, "layout", expenses)
}

func (h *ExpenseHandler) AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	newExpense := models.Expense{}
	newExpense.Description = r.FormValue("description")
	var err error
	num, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	newExpense.Amount = num
	newExpense.Date = r.FormValue("date")
	err = h.Repo.Create(&newExpense)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
