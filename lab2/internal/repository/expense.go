package repository

import (
	"database/sql"
	"kgu_server/internal/models"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := r.db.Query("SELECT id, amount, description, date FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var e models.Expense
		err = rows.Scan(&e.Id, &e.Amount, &e.Description, &e.Date)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *ExpenseRepository) Create(expense *models.Expense) error {
	_, err := r.db.Exec("INSERT INTO expenses (amount, description, date) VALUES (?, ?, ?)",
		expense.Amount, expense.Description, expense.Date)
	if err != nil {
		return err
	} else {
		return nil
	}
}
