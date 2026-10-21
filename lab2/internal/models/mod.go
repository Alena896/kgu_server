package models

type Expense struct {
	Id          int64
	Amount      float64
	Description string
	Date        string
	Category    string
	CategoryId  int64
	Action      string
}
