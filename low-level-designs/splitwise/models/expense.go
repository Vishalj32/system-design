package models

// ExpenseType is the type of the expense (EQUAL, EXACT, PERCENT)
type ExpenseType string

const (
	EQUAL   ExpenseType = "EQUAL"
	EXACT   ExpenseType = "EXACT"
	PERCENT ExpenseType = "PERCENT"
)

// Expense represents an expense in the application
type Expense struct {
	PayerId       string
	Amount        float64
	UsersInvolved []string
	Splits        []float64
	ExpenseType   ExpenseType
}
