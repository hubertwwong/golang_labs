package main

import (
  "encoding/json"
  "fmt"
)

// Marshal. Takes a go type to json.
func encode01() {
  bolB, _ := json.Marshal(true)
  fmt.Println(string(bolB))

  intB, _ := json.Marshal(1)
  fmt.Println(string(intB))

  floatB, _ := json.Marshal(2.34)
  fmt.Println(string(floatB))

  strB, _ := json.Marshal("gopher")
  fmt.Println(string(strB))

  slc := []string{"apple", "peach"}
  sliceB, _ := json.Marshal(slc)
  fmt.Println(string(sliceB))

  maps := map[string]int{"apple": 5, "orange": 20}
  mapsB, _ := json.Marshal(maps)
  fmt.Println(string(mapsB))
}

func main() {
  encode01()
}

/*

marshalling.
takes a golang type and converts it to a json type

https://gobyexample.com/json
https://gobyexample.com/worker-pools


*/