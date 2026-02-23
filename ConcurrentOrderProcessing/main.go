package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount float64
	Status string
}

type Formatter interface {
	print() string
}

func (order Order) print() string {
	return fmt.Sprintf("Order ID: %d, Amount: %.2f, Status: %s", order.ID, order.Amount, order.Status)
}
func processOrder(ctx context.Context, id int, amount float64, ch chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a timer but don't block yet
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-ctx.Done():
		// The signal was sent to stop!
		fmt.Printf("Worker %d: received cancellation, stopping...\n", id)
		return // Exit the function early
	case <-timer.C:
		// The work is done
		order := Order{ID: id, Amount: amount, Status: ""}
		if amount > 100 {
			order.Status = "High Value - Requires Inspection"
		} else {
			order.Status = "Approved"
		}
		ch <- order
	}
}

func main() {
	// Create a context that cancels in 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Good practice to clean up

	orderList := []float64{45.50, 720.00, 10.0, 561}
	chResult := make(chan Order, len(orderList))
	var wg sync.WaitGroup

	for _, amt := range orderList {
		wg.Add(1)
		// Pass ctx to the worker
		go processOrder(ctx, rand.Intn(1000), amt, chResult, &wg)
	}

	// Standard background waiter
	go func() {
		wg.Wait()
		close(chResult)
	}()

	for {
		select {
		case order, ok := <-chResult:
			if !ok {
				fmt.Println("All done!")
				return
			}
			fmt.Println(order.print())
		case <-ctx.Done():
			// This triggers when the 3s are up
			fmt.Println("Main: Timeout reached, exiting.")
			// We wait a tiny bit just to see the workers log their exit
			time.Sleep(100 * time.Millisecond)
			return
		}
	}
}
