# Concurrent Order Processor in Go

A high-performance order processing simulation built to demonstrate **Go's Concurrency Model**. This project showcases how to handle multiple tasks simultaneously using idiomatic Go patterns like Goroutines, Channels, and WaitGroups.



## 🚀 Key Features

* **Goroutines**: Processes multiple orders in parallel instead of sequentially.
* **Channels**: Implements safe communication between workers and the main thread.
* **WaitGroups**: Synchronizes the lifecycle of background tasks to prevent race conditions.
* **The "Background Waiter" Pattern**: Uses a dedicated goroutine to manage channel closure, preventing deadlocks.
* **Custom Interfaces**: Implements a `Formatter` interface for decoupled and professional logging.

## 🛠️ Architecture

The system follows a producer-consumer flow where the main thread spawns workers and listens for results through a centralized communication pipe.



## 📋 Logic Flow

1.  **Input**: A slice of floating-point order amounts.
2.  **Concurrency**: For every amount, a Goroutine is spawned to handle `processOrder`.
3.  **Validation**:
    * Orders > **$100.00** are flagged as `"High Value - Requires Inspection"`.
    * Orders ≤ **$100.00** are marked as `"Approved"`.
4.  **Synchronization**: A `sync.WaitGroup` tracks all active workers.
5.  **Clean Exit**: Once all workers finish, the channel is closed, allowing the results loop to terminate gracefully.

## 💻 Sample Output

```text
Processing Orders...
Order ID: 456 | Amount: $720.00 | Status: High Value - Requires Inspection
Order ID: 123 | Amount: $45.50  | Status: Approved
Order ID: 882 | Amount: $600.00 | Status: High Value - Requires Inspection
Order ID: 015 | Amount: $15.00  |