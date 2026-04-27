package storage

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"os"
)

func Load() ([]models.Expense, error) {
	data, err := os.ReadFile("expenses.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Expense{}, nil
		} else {
			return nil, err
		}
	}

	if len(data) == 0 {
		return []models.Expense{}, nil
	}

	var expenses []models.Expense
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func Save(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile("expenses.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
