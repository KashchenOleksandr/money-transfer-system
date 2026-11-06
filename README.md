# 💸 Money Transfer System

A simple money transfer system written in Go.

The project demonstrates the use of structs, methods, error handling, goroutines, channels, WaitGroups, and Mutex for concurrent transaction processing.

## 🚀 Features

- User management
- Deposit and withdraw operations
- Transaction processing between users
- Input validation and error handling
- Concurrent transaction processing using goroutines
- Buffered channels for communication
- WaitGroup for synchronization
- Mutex to protect user balances from concurrent access

## 📂 Project Versions

The project is divided into four Git branches:

### v1 — Basic User Operations

- Created the `User` structure.
- Implemented `Deposit()` and `Withdraw()` methods.
- Added balance validation and error handling.

### v2 — Transaction System

- Created the `Transaction` structure.
- Implemented the `PaymentSystem` structure.
- Added user and transaction management.
- Implemented transaction processing.

### v3 — Concurrent Transaction Processing

- Added goroutines and workers.
- Implemented buffered channels.
- Used `sync.WaitGroup` to wait for all workers.
- Added concurrent transaction processing.

### v4 — Thread-Safe Balance Operations

- Added `sync.Mutex` to the `User` structure.
- Protected `Deposit()` and `Withdraw()` operations.
- Used the race detector to test concurrent access.

## 🛠️ Requirements

- Go 1.26.6 or newer

## ▶️ How to Run

Clone the repository:

```bash
git clone https://github.com/KashchenOleksandr/money-transfer-system.git
```

Navigate to the project directory:

```bash
cd money-transfer-system
```

Run the application:

```bash
go run .
```

Run with the race detector:

```bash
go run -race .
```