package main

import "fmt"
import "time"

func main() {
  i := 2

  fmt.Println("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one");
  case 2:
    fmt.Println("two");
  case 3:
    fmt.Println("three");
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("its the weekend");
  default:
    fmt.Println("its the weekday");
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("before noon")
  default:
    fmt.Println("after noon")
  }
}
