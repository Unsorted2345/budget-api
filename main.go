package main

import "github.com/gin-gonic/gin"

func main() {
	initDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" { c.AbortWithStatus(204); return }
		c.Next()
	})
	r.GET("/budgets", getBudgets)
	r.GET("/transactions", getTransactions)
	r.GET("budgets/:id", getBudget)
	r.GET("transactions/:id", getTransaction)
	r.GET("/report", getReport)
	r.POST("/budgets", setBudgets)
	r.POST("/transactions", addTransaction)
	r.DELETE("/budgets/:id", deleteBudget)
	r.DELETE("/transactions/:id", deleteTransaction)
	r.PUT("/budgets/:id", updateBudget)
	r.PUT("/transactions/:id", updateTransaction)
	r.Run(":6090")
}
