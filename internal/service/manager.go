package service

import (
	"expense-tracker/internal/models"
	"expense-tracker/internal/storage"
	"fmt"
	"strconv"
	"time"
)

func AddExpense(description string, amount string) error {
	expenses, err := storage.Load()
	if err != nil {
		return err
	}

	maxID := 0
	for _, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	newID := maxID + 1

	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d-%02d-%02d", year, month, day)

	newExpense := models.Expense{
		ID:          newID,
		Description: description,
		Amount:      amount,
		Date:        date,
	}

	expenses = append(expenses, newExpense)

	return storage.Save(expenses)
}

func UpdateExpense(ID int, newDescription string, newAmount string) error {
	expenses, err := storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i := range expenses {
		if expenses[i].ID == ID {
			expenses[i].Description = newDescription
			expenses[i].Amount = newAmount
			found = true
			break
		}
	}

	if found == false {
		return fmt.Errorf("expense with ID: %v is not found", ID)
	}

	return storage.Save(expenses)
}

func DeleteExpense(ID int) error {
	expenses, err := storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i := range expenses {
		if expenses[i].ID == ID {
			found = true
			expenses = append(expenses[:i], expenses[i+1:]...)
			break
		}
	}

	if found == false {
		return fmt.Errorf("expense with ID: %v is not found", ID)
	}

	return storage.Save(expenses)
}

func ListExpenses() error {
	expenses, err := storage.Load()
	if err != nil {
		return err
	}

	fmt.Println("ID  Date       Description  Amount")
	for _, e := range expenses {
		amount := "$" + e.Amount
		fmt.Printf("%v   %v  %v       %v\n", e.ID, e.Date, e.Description, amount)
	}

	return nil
}

func Sumary() (int, error) {
	expenses, err := storage.Load()
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, e := range expenses {
		amount, err := strconv.Atoi(e.Amount)
		if err != nil {
			return 0, err
		}
		sum += amount
	}

	return sum, nil
}

func SummaryForMonth(month int) (int, error) {
	expenses, err := storage.Load()
	if err != nil {
		return 0, err
	}

	sum := 0
	monthExp := 0
	var amount int

	for _, e := range expenses {
		monthExp, err = strconv.Atoi(e.Date[5:7])
		if err != nil {
			return 0, err
		}
		if monthExp == month {
			amount, err = strconv.Atoi(e.Amount)
			if err != nil {
				return 0, err
			}
			sum += amount
		}
	}

	return sum, nil
}
