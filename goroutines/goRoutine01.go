package main

import (
  "fmt"
  "time"
)

func f(from string) {
  for i := 0 ; i < 3 ; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  f("direct")

  go f("goroutine")

  // go routine on anon func...
  go func(msg string) {
    fmt.Println(msg)
  }("going")

  // adding a sleep timer.
  // This will give the goroutines a chance to finish...
  // innstead of user input.
  time.Sleep(time.Second * 2)

  // user input. I think this gives your go routine to finish.
  //var input string
  //fmt.Scanln(&input)
  //fmt.Println("done")
}

/*

https://gobyexample.com/goroutines
https://gobyexample.com/timeouts



*/