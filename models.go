package main

import "time"

type Transaction struct {
	ID          int        `json:"id"`
	Type        string     `json:"type"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description,omitempty"`
	Date        *time.Time `json:"date"`
}

type Budget struct {
	ID     int        `json:"id"`
	Amount float64    `json:"amount"`
	Month  time.Month `json:"month"`
	Year   int        `json:"year"`
}

type Report struct {
	Month     time.Month `json:"month"`
	Year      int        `json:"year"`
	Budget    float64    `json:"budget"`
	Expenses  float64    `json:"expenses"`
	Income    float64    `json:"income"`
	Remaining float64    `json:"remaining"`
	Change    float64    `json:"change"`
}
