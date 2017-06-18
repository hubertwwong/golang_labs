package main

import "fmt"

func ping(pings chan<- string, msg string) {
  pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)

  // sends a message to ping channel
  ping(pings, "passed message")
  // passes message from ping to pong channel
  pong(pings, pongs)
  // output pong channel..
  fmt.Println(<- pongs)
}