package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  // non blocking receive.
  select {
    case msg := <- messages:
      fmt.Println("rec msg", msg)
    default:
      fmt.Println("no msg rec")
  }

  msg := "hi"

  // non blocking send.
  select {
    case messages <- msg:
      fmt.Println("sent msg", msg)
    default:
      fmt.Println("no msg sent")
  }

  select {
    case msg := <- messages:
      fmt.Println("rec msg", msg)
    case sig := <- signals:
      fmt.Println("rec signal", sig)
    default:
      fmt.Println("no act")
  }
}