package models

type Expense struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
}

