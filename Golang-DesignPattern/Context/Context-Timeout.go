package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed successfully")
    case <-ctx.Done():
        fmt.Println("Operation timed out")
    }
}
