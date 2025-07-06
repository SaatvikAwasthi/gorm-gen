# GORM Gen Example Project

This project demonstrates how to use [GORM Gen](https://gorm.io/gen/) for type-safe, code-generated database access in Go. It features basic CRUD operations, transactions, and model relationships using an in-memory SQLite database.

## Features

- Type-safe database queries with GORM Gen
- Auto code generation for models and queries
- Example CRUD operations and transactions
- Model relationships (User has many Products)
- In-memory SQLite for easy setup

## Project Structure

- `cmd/main.go` — Entry point, runs code generation and service logic
- `repository/` — Database connection and GORM Gen code generation logic
- `dao/` — Data models (`User`, `Product`)
- `contract/` — Query interface for custom SQL
- `service/` — Example service using generated code
- `generated/` — Auto-generated code (after running the app)

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

```sh
go get -u gorm.io/gorm gorm.io/driver/sqlite gorm.io/gen
go mod tidy
```

### Usage

Run the application (this will generate code and execute the service):

```sh
go run cmd/main.go
```

You should see logs for code generation and CRUD operations.

## How It Works

- Models are defined in `dao/`
- `repository/base.go` sets up an in-memory SQLite DB and runs GORM Gen to generate type-safe code
- `service/service.go` demonstrates CRUD, transactions, and preloading relationships using the generated code
