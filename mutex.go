// bank.go
package main

import "sync"
import "time"
import "fmt"

func main() {
    alice := 10000
    bob := 10000
    var mu sync.Mutex

    total := alice + bob

    go func() {
        for i := 0; i < 1000; i++ {
            mu.Lock()
      // defer mu.Unlock()
            alice -= 1
            bob += 1
            mu.Unlock()
        }
    }()
    go func() {
        for i := 0; i < 1000; i++ {
            mu.Lock()
      // defer mu.Unlock()
            bob -= 1
            alice += 1
            mu.Unlock()
        }
    }()

    start := time.Now()
    for time.Since(start) < 1*time.Second {
        mu.Lock()
        if alice+bob != total {
            fmt.Printf("observed violation, alice = %v, bob = %v, sum = %v\n", alice, bob, alice+bob)
        }
        mu.Unlock()
    }
  fmt.Printf("alie+bob= %v\n", alice+bob)
}