package main

import "sync"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())
	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
  
	for i := 0; i < 10; i++ {
	  go func() {
		vote := requestVote()
		mu.Lock()
		defer mu.Unlock() // remember to unlock the mutex after broadcasting
		if vote {
		  count++
		}
		finished++
		// broadcast to the threads that waitting on the condition variable.
		cond.Broadcast()
	  }()
	}
	mu.Lock()   // the lock here also avoid the lost-wakeup issue.
	for count < 5 && finished != 10 {
	  // this will block the current thread and give up its lock atomically 
	  // so that other threads can make progress.
	  cond.Wait()
	  // if someone wakes this up. the current thread will try to grab the
	  // lock again. and then check the count and finished while holding 
	  // the lock
	}
	if count >= 5 {
	  println("received 5+ votes!")
	} else {
	  println("lost")
	}
	mu.Unlock()
  }
  func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int() % 2 == 0
  }