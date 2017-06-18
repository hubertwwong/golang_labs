package main

import "fmt"

func main() {
  messages := make(chan string)

  // using an anon func
  // sending a ping string to a messange channel
  // goroutine too.
  go func() {messages <- "ping"}()

  // can you do this?
  // why does this cause a deadlock.
  // messages <- "ping2"
  
  msg := <-messages
  fmt.Println(msg)
}