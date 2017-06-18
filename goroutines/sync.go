package main

import (
  "fmt"
  "time"
)

// This prints working and then done.
// then sends a boolen to the done channel?
func worker(done chan bool) {
  fmt.Println("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)

  // this is the sync.
  // you don't acutally care what is returned.
  // go has a mechanic that blocks until you get a message on a channel.
  <-done
}