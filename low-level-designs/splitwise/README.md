# Expense Sharing Application

This is a simple expense-sharing application built using Go. The application allows users to split expenses among a group of people and keep track of balances between them. 

## Features
- **Expense Types**: 
  - `EQUAL`: Splits the expense equally among all users.
  - `EXACT`: Splits the expense based on exact amounts provided for each user.
  - `PERCENT`: Splits the expense based on the percentage share provided for each user.
  
- **Balance Management**: 
  - Keeps track of who owes how much to whom.
  - Allows viewing of balances between any user and others involved in their expenses.
  
- **Customizable**: Can be extended to add expense names, notes, or images. 

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) 1.16 or higher.

### How to Run

1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd expense-sharing-application
    ```

2. Build and run the Go application:
    ```bash
    go run main.go
    ```

### Usage

The application tracks expenses in three ways: **EQUAL**, **EXACT**, and **PERCENT**. You can add expenses, and the system will manage the balances between users.

### Input Format

The following types of input can be processed by the application:

1. **Expense**: Add an expense.
    ```
    EXPENSE <user-id-who-paid> <total-amount> <no-of-users> <user-ids-involved> <EQUAL/EXACT/PERCENT> <amounts/percentages-if-needed>
    ```

   - **EQUAL**: Splits the amount equally.
   - **EXACT**: Splits based on exact amounts provided for each user.
   - **PERCENT**: Splits based on percentage values for each user (sum must equal 100%).

2. **Show Balances for All Users**:
    ```
    SHOW
    ```

3. **Show Balances for a Specific User**:
    ```
    SHOW <user-id>
    ```

### Examples

#### EQUAL Expense
User1 pays Rs. 1000, split equally among User1, User2, User3, and User4:
EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL


#### EXACT Expense
User1 pays Rs. 1250 for User2 and User3. Exact amounts are specified:
EXPENSE u1 1250 2 u2 u3 EXACT 370 880


#### PERCENT Expense
User4 pays Rs. 1200, split by percentage between User1, User2, User3, and User4:
EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20


#### Show Balances for All Users
Display balances for all users:
SHOW

#### Show Balances for a Specific User
Display balances for User1:
SHOW u1


### Sample Output
Balances for u1: u2 owes u1: 620.00 u3 owes u1: 1130.00 u1 owes u4: 230.00

All balances: u2 owes u1: 620.00 u3 owes u1: 1130.00 u1 owes u4: 230.00 u2 owes u4: 240.00 u3 owes u4: 240.00
