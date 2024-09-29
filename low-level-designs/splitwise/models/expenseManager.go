package models

import (
	"fmt"
	"math"
)

// ExpenseManager handles all the users, balances and expenses
type ExpenseManager struct {
	Users    map[string]*User
	Balances map[string]map[string]float64
}

// NewExpenseManager creates a new instance of ExpenseManager
func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		Users:    make(map[string]*User, 0),
		Balances: make(map[string]map[string]float64),
	}
}

// AddUser adds a new user
func (em *ExpenseManager) AddUser(user *User) {
	em.Users[user.ID] = user
	em.Balances[user.ID] = make(map[string]float64)
}

func (em *ExpenseManager) updateBalance(payerId, userId string, amount float64) {
	em.Balances[userId][payerId] += amount
}

// splitEqual handles EQUAL split of expenses
func (em *ExpenseManager) splitEqual(expense *Expense) {
	totalUsers := len(expense.UsersInvolved)
	equalSplit := math.Round((expense.Amount/float64(totalUsers))*100) / 100

	for _, userId := range expense.UsersInvolved {
		if userId == expense.PayerId {
			em.updateBalance(expense.PayerId, userId, equalSplit)
		}
	}
}

// splitExact handles EXACT split of expense
func (em *ExpenseManager) splitExact(expense *Expense) {
	sum := 0.0
	for _, split := range expense.Splits {
		sum += split
	}

	if math.Round(sum*100)/100 != expense.Amount {
		fmt.Println("Error: Exact split amounts do not match the total amount")
		return
	}

	for idx, userId := range expense.UsersInvolved {
		if userId != expense.PayerId {
			em.updateBalance(expense.PayerId, userId, expense.Splits[idx])
		}
	}
}

// splitPercent handles PERCENT split of expense
func (em *ExpenseManager) splitPercent(expense *Expense) {
	totalPercent := 0.0
	for _, percent := range expense.Splits {
		totalPercent += percent
	}

	if math.Round(totalPercent*100)/100 != 100.0 {
		fmt.Println("Error: Total percentage is not 100")
		return
	}

	for idx, userId := range expense.UsersInvolved {
		if userId != expense.PayerId {
			share := math.Round((expense.Amount*(expense.Splits[idx]/100))*100) / 100
			em.updateBalance(expense.PayerId, userId, share)
		}
	}

}

// AddExpense adds a new expense to system
func (em *ExpenseManager) AddExpense(expense *Expense) {
	switch expense.ExpenseType {
	case EQUAL:
		{
			em.splitEqual(expense)
		}
	case EXACT:
		{
			em.splitExact(expense)
		}
	case PERCENT:
		{
			em.splitPercent(expense)
		}
	}
}

// ShowBalance displays balance of the specific userID
func (em *ExpenseManager) ShowBalance(userId string) {
	found := false
	for otherUserid, balance := range em.Balances[userId] {
		if balance != 0 {
			if balance > 0 {
				fmt.Printf("%s owes %s: %.2f\n", userId, otherUserid, balance)
			} else {
				fmt.Printf("%s owes %s: %.2f\n", otherUserid, userId, -balance)
			}
			found = true
		}
	}

	if !found {
		fmt.Println("No Balances")
	}
}

// ShowAllBalances displays all the user balance in the system
func (em *ExpenseManager) ShowAllBalances() {
	for userId, userBalances := range em.Balances {
		for otherUserId, balance := range userBalances {
			if balance > 0 {
				fmt.Printf("%s owes %s: %.2f\n", userId, otherUserId, balance)
			} else {
				fmt.Printf("%s owes %s: %.2f\n", otherUserId, userId, -balance)
			}
		}
	}
}
