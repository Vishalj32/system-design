package main

import (
	"fmt"

	"github.com/Vishalj32/system-design/low-level-designs/splitwise/models"
)

func main() {
	// Initialize the expense manager
	em := models.NewExpenseManager()

	// Add users
	em.AddUser(&models.User{ID: "u1", Name: "User1", Email: "user1@example.com", Phone: "1234567890"})
	em.AddUser(&models.User{ID: "u2", Name: "User2", Email: "user2@example.com", Phone: "1234567891"})
	em.AddUser(&models.User{ID: "u3", Name: "User3", Email: "user3@example.com", Phone: "1234567892"})
	em.AddUser(&models.User{ID: "u4", Name: "User4", Email: "user4@example.com", Phone: "1234567893"})

	// Add expenses
	em.AddExpense(&models.Expense{
		PayerId:       "u1",
		Amount:        1000,
		UsersInvolved: []string{"u1", "u2", "u3", "u4"},
		ExpenseType:   models.EQUAL,
	})

	em.AddExpense(&models.Expense{
		PayerId:       "u1",
		Amount:        1250,
		UsersInvolved: []string{"u2", "u3"},
		ExpenseType:   models.EXACT,
		Splits:        []float64{370, 880},
	})

	em.AddExpense(&models.Expense{
		PayerId:       "u4",
		Amount:        1200,
		UsersInvolved: []string{"u1", "u2", "u3", "u4"},
		ExpenseType:   models.PERCENT,
		Splits:        []float64{40, 20, 20, 20},
	})

	// Show balances
	fmt.Println("Balances for u1:")
	em.ShowBalance("u1")

	fmt.Println("\nAll balances:")
	em.ShowAllBalances()
}
