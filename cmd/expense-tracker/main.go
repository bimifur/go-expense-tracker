package main

import (
	"expense-tracker/internal/service"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Доступные комманды:\nadd [описание] [расход] - добавить расход\ndelete [ID] - удалить расход\nupdate [ID] [новое описание] [новый расход] - обновить расход\nlist - список всех расходов\nsumary - сумма всех расходов")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("need: add [описание] [расход]")
			return
		}

		description := os.Args[2]
		amount := os.Args[3]

		err := service.AddExpense(description, amount)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Задача успешно добавлена!")

	case "delete":
		if len(os.Args) < 2 {
			fmt.Println("need: delete [ID]")
			return
		}

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		err = service.DeleteExpense(ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Задача успешно удалена!")

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("need: update [ID] [новое описание] [новый расход]")
			return
		}

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		newDescription := os.Args[3]
		newAmount := os.Args[4]

		err = service.UpdateExpense(ID, newDescription, newAmount)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("Задача обновлена успешно!")

	case "list":
		err := service.ListExpenses()
		if err != nil {
			fmt.Println("error:", err)
		}

	case "summary":
		if len(os.Args) == 2 {
			summary, err := service.Sumary()
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			fmt.Printf("Сумма всех трат: $%v", summary)
		} else if len(os.Args) == 3 {
			month, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			summary, err := service.SummaryForMonth(month)
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			fmt.Printf("Сумма трат за %v месяц: $%v", month, summary)
		}
	}

}
