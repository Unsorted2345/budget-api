# Budget API

A simple REST API for tracking budgets and transactions, built with Go, Gin and SQLite.

## Features

- Set monthly budgets
- Track income and expenses
- Generate monthly reports (budget, income, expenses, remaining, balance)
- Filter transactions by month and year

## Tech Stack

- **Go** – Backend language
- **Gin** – HTTP framework
- **SQLite** – Database

## Getting Started

### Prerequisites

- Go 1.20+
- GCC (required for go-sqlite3)

### Installation

```bash
git clone https://github.com/your-username/budget-api.git
cd budget-api
go mod tidy
go run .
```

The server starts on `http://localhost:8080`.

## API Endpoints

### Budgets

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/budgets` | Get all budgets |
| GET | `/budgets/:id` | Get a single budget |
| POST | `/budgets` | Create a budget |
| PUT | `/budgets/:id` | Update a budget |
| DELETE | `/budgets/:id` | Delete a budget |

#### Budget Object

```json
{
  "id": 1,
  "amount": 1500.00,
  "month": 5,
  "year": 2026
}
```

> Only one budget per month/year is allowed.

---

### Transactions

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/transactions` | Get all transactions |
| GET | `/transactions/:id` | Get a single transaction |
| POST | `/transactions` | Create a transaction |
| PUT | `/transactions/:id` | Update a transaction |
| DELETE | `/transactions/:id` | Delete a transaction |

#### Transaction Object

```json
{
  "id": 1,
  "type": "expense",
  "amount": 49.99,
  "description": "Groceries",
  "date": "2026-05-14T10:00:00Z"
}
```

- `type` must be either `"income"` or `"expense"`
- `date` is optional — defaults to the current timestamp if not provided

---

### Report

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/report?month=5&year=2026` | Get monthly report |

#### Report Response

```json
{
  "month": 5,
  "year": 2026,
  "budget": 1500.00,
  "income": 200.00,
  "expenses": 340.00,
  "remaining": 1160.00,
  "change": -140.00
}
```

- `remaining` = budget − expenses
- `change` = income − expenses

## Project Structure

```
budget-api/
├── main.go       # Server setup and routes
├── db.go         # Database connection and initialization
├── models.go     # Data structures
└── handlers.go   # HTTP handlers
```

## License

MIT
