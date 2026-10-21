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

func (r *ExpenseRepository) GetAll(category string) ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := r.db.Query(`SELECT expenses.id, expenses.amount, expenses.description, expenses.date, categories.name_category 
							FROM expenses INNER JOIN categories ON expenses.category_id = categories.id`)
	if category != "" {
		rows, err = r.db.Query(`SELECT expenses.id, expenses.amount, expenses.description, expenses.date, categories.name_category 
		FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE categories.name_category = ?`, category)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var e models.Expense
		err = rows.Scan(&e.Id, &e.Amount, &e.Description, &e.Date, &e.Category)
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

func (r *ExpenseRepository) Create(expense *models.Expense, category string) error {
	var catId int
	err := r.db.QueryRow("SELECT id FROM categories WHERE name_category = ?", category).Scan(&catId)
	if err == sql.ErrNoRows {
		res, _ := r.db.Exec("INSERT INTO categories (name_category) VALUES (?)", category)
		insertId, _ := res.LastInsertId()
		catId = int(insertId)
	}
	_, err = r.db.Exec("INSERT INTO expenses (amount, description, date, category_id) VALUES (?, ?, ?, ?)",
		expense.Amount, expense.Description, expense.Date, catId)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r *ExpenseRepository) GetById(id int) (*models.Expense, error) {
	var e models.Expense
	err := r.db.QueryRow("SELECT id, amount, description, date FROM expenses WHERE id = ?", id).
		Scan(&e.Id, &e.Amount, &e.Description, &e.Date)
	if err != nil {
		return nil, err
	} else {
		return &e, nil
	}
}

func (r *ExpenseRepository) Update(expense *models.Expense) error {
	var ctgrId int64
	err := r.db.QueryRow("SELECT id FROM categories WHERE name_category = ?", expense.Category).Scan(&ctgrId)
	if err == sql.ErrNoRows {
		res, _ := r.db.Exec("INSERT INTO categories (name_category) VALUES (?)", expense.Category)
		ctgrId, _ = res.LastInsertId()
	}
	_, err = r.db.Exec("UPDATE expenses SET amount = ?, description = ?, date = ?, category_id = ? WHERE id = ?",
		expense.Amount, expense.Description, expense.Date, ctgrId, expense.Id)
	return err
}

func (r *ExpenseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM expenses WHERE id = ?", id)
	return err
}
