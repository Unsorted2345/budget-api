package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getBudgets(c *gin.Context) {
	var budgets []Budget
	rows, err := db.Query("SELECT id, amount, month, year FROM budgets")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for rows.Next() {
		var budget Budget
		err := rows.Scan(&budget.ID, &budget.Amount, &budget.Month, &budget.Year)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		budgets = append(budgets, budget)
	}
	c.JSON(200, budgets)
}

func getBudget(c *gin.Context) {
	var budget Budget
	id := c.Param("id")
	var err error

	row := db.QueryRow("SELECT id, amount, month, year FROM budgets WHERE id = ?", id)

	err = row.Scan(&budget.ID, &budget.Amount, &budget.Month, &budget.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Budget with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, budget)
}

func setBudgets(c *gin.Context) {
	var budget Budget
	var err error
	err = c.ShouldBindJSON(&budget)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := db.Exec("INSERT INTO budgets(amount, month, year) VALUES (?, ?, ?)", budget.Amount, budget.Month, budget.Year)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	budget.ID = int(lastID)
	c.JSON(201, budget)
}

func updateBudget(c *gin.Context) {
	var budget Budget
	id := c.Param("id")
	var rowsAffected int64
	var err error

	row := db.QueryRow("SELECT id, amount, month, year FROM budgets WHERE id = ?", id)

	err = row.Scan(&budget.ID, &budget.Amount, &budget.Month, &budget.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Budget with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindJSON(&budget)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("UPDATE budgets SET amount = ?, month = ?, year = ? WHERE id = ?", budget.Amount, budget.Month, budget.Year, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(404, fmt.Sprintf("Budget with ID %s doesn't exist", id))
		return
	}
	c.JSON(200, budget)
}

func deleteBudget(c *gin.Context) {
	id := c.Param("id")
	var budget Budget
	var err error
	row := db.QueryRow("SELECT id, amount, month, year FROM budgets WHERE id = ?", id)

	err = row.Scan(&budget.ID, &budget.Amount, &budget.Month, &budget.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Budget with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("DELETE FROM budgets WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

func getTransactions(c *gin.Context) {
	var transactions []Transaction
	rows, err := db.Query("SELECT id, type, amount, description, date FROM transactions")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Description, &transaction.Date)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, transaction)
	}
	c.JSON(200, transactions)
}

func getTransaction(c *gin.Context) {
	var transaction Transaction
	id := c.Param("id")
	var err error

	row := db.QueryRow("SELECT id, type, amount, description, date FROM transactions WHERE id = ?", id)

	err = row.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Description, &transaction.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Transaction with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transaction)
}

func addTransaction(c *gin.Context) {
	var transaction Transaction
	var err error
	err = c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if transaction.Date == nil {
		now := time.Now()
		transaction.Date = &now
	}
	result, err := db.Exec("INSERT INTO transactions(type, amount, description, date) VALUES (?, ?, ?, ?)", transaction.Type, transaction.Amount, transaction.Description, transaction.Date)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	transaction.ID = int(lastID)
	c.JSON(201, transaction)
}

func updateTransaction(c *gin.Context) {
	var transaction Transaction
	id := c.Param("id")
	var rowsAffected int64
	var err error

	row := db.QueryRow("SELECT id, type, amount, description, date FROM transactions WHERE id = ?", id)

	err = row.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Description, &transaction.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Transaction with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("UPDATE transactions SET type = ?, amount = ?, description = ?, date = ? WHERE id = ?", transaction.Type, transaction.Amount, transaction.Description, transaction.Date, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(404, fmt.Sprintf("Transaction with ID %s doesn't exist", id))
		return
	}
	c.JSON(200, transaction)
}

func deleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction Transaction
	var err error

	row := db.QueryRow("SELECT id, type, amount, description, date FROM transactions WHERE id = ?", id)

	err = row.Scan(&transaction.ID, &transaction.Type, &transaction.Amount, &transaction.Description, &transaction.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, fmt.Sprintf("Transaction with ID %s doesn't exist", id))
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

func getReport(c *gin.Context) {
	var report Report
	month := c.Query("month")
	year := c.Query("year")
	var err error

	MonthInt, err := strconv.Atoi(month)
	if err != nil {
		c.JSON(500, fmt.Sprintf("Error in `month`: Invalid value `%s`", month))
		return
	}
	report.Month = time.Month(MonthInt)

	db.QueryRow("SELECT SUM(amount) FROM transactions WHERE type = 'expense' AND strftime('%m', date) = ? AND strftime('%Y', date) = ?", fmt.Sprintf("%02d", MonthInt), year).Scan(&report.Expenses)
	db.QueryRow("SELECT SUM(amount) FROM transactions WHERE type = 'income' AND strftime('%m', date) = ? AND strftime('%Y', date) = ?", fmt.Sprintf("%02d", MonthInt), year).Scan(&report.Income)
	db.QueryRow("SELECT amount FROM budgets WHERE month = ? AND year = ?", month, year).Scan(&report.Budget)

	report.Year, err = strconv.Atoi(year)
	if err != nil {
		c.JSON(500, fmt.Sprintf("Error in `year`: Invalid value `%s`", year))
		return
	}

	report.Remaining = report.Budget - report.Expenses
	report.Change = report.Income - report.Expenses

	c.JSON(200, report)
}
