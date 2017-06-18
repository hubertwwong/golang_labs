package main

import (
  "encoding/json" 
  "fmt"
  "os"
)

type Response1 struct {
  Page int
  Fruits []string
}

type Response2 struct {
  Page int `json:"page"`
  Fruits []string `json:"fruits"`
}

func decode01() {
  fmt.Println("> deccode01")
  // marshalling a struct
  res1D := &Response1{
      Page: 1, 
      Fruits: []string{"apple", "peach", "pear"}}

  res1B, _ := json.Marshal(res1D)
  fmt.Println(string(res1B))

  // marshalling using the "json" types.
  res2D := &Response2{
    Page: 1,
    Fruits: []string{"apple", "orange"}}
  res2B, _ := json.Marshal(res2D)
  fmt.Println(string(res2B))

  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

  // This is the generic interface to decode json.
  // You have to cast data.
  var dat map[string]interface{}

  // this is the decoding of byt...
  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)
  
  // This is the casting from interface to byte
  num := dat["num"].(float64)
  num += 1.0
  fmt.Println("> basic decode", num)

  // nested decode example....
  // this is off the byt.
  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println("> nested example",str1)
}

func decode02() {
  fmt.Println("\n> decode02 > ")
  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple": 5, "lettuce": 7}
  enc.Encode(d)
}

func decode03() {
  fmt.Println("\n> decode03 > unmarhal to a type. This removes the casting requirements for fields.")
  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res := Response2{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println(res)
  fmt.Println(res.Fruits[0])
}



func main() {
  decode01()
  decode02()
  decode03()
  fmt.Println("> done")
}

/*
On struct you can use a colon to id the field...
Im guessing you can use this to do out of order.

*/