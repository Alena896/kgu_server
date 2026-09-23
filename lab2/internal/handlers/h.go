package handlers

import (
	"fmt"
	"html/template"
	"kgu_server/internal/models"
	"net/http"
	"strconv"
)

var ExpensesList []models.Expense

func ListExpensesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, l := template.ParseFiles("web/templates/layout.html", "web/templates/index.html")
	fmt.Println(l)
	tmpl.ExecuteTemplate(w, "layout", ExpensesList)
}

func AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newExpense := models.Expense{}
	newExpense.Description = r.FormValue("description")
	num, _ := strconv.ParseInt(r.FormValue("amount"), 10, 64)
	newExpense.Amount = num
	newExpense.Date = r.FormValue("date")
	ExpensesList = append(ExpensesList, newExpense)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
