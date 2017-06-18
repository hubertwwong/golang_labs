package main

import (
  "fmt"
  "time"
)

// This wont finish since we set the time to 2
// should print timeout.
func t1() {
  c1 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c1 <- "res 1"
  }()

  select {
    case res := <- c1:
      fmt.Println(res)
    case <- time.After(time.Second * 1):
      fmt.Println("timeout 1")
  }
}

// Since the timeout is longer,
// this will print res 2.
func t2() {
  c2 := make(chan string, 1)
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "res 2"
  }()

  select {
    case res := <-c2:
      fmt.Println(res)
    case <- time.After(time.Second * 3):
      fmt.Println("timeout 2")
  }
}

func main() {
  t1()
  t2()
}