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
	categoreParam := r.URL.Query().Get("category")
	expenses, err := h.Repo.GetAll(categoreParam)
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
	categoryInput := r.FormValue("name_category")
	err = h.Repo.Create(&newExpense, categoryInput)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExpenseHandler) DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	idNum, _ := strconv.Atoi(idParam)
	err := h.Repo.Delete(idNum)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ExpenseHandler) EditExpenseHandler(w http.ResponseWriter, r *http.Request) {
	idEdit := r.FormValue("id")
	idInt, _ := strconv.Atoi(idEdit)
	var updatedExpense models.Expense
	if r.Method == http.MethodPost {
		updatedExpense.Id = int64(idInt)
		idAmount, _ := strconv.Atoi(r.FormValue("amount"))
		updatedExpense.Amount = float64(idAmount)
		updatedExpense.Description = r.FormValue("description")
		updatedExpense.Date = r.FormValue("date")
		updatedExpense.Category = r.FormValue("name_category")
		err := h.Repo.Update(&updatedExpense)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	expense, err := h.Repo.GetById(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, _ := template.ParseFiles("web/templates/layout.html", "web/templates/edit.html")
	tmpl.ExecuteTemplate(w, "layout", expense)
}

//
