package main

import (
  "encoding/json"
  "fmt"
)

type ColorfulEcosystem struct {
  Things []ColoredThings
}

type ColoredThings interface {
  Color() string
}

// ===== PLANT ======

type Plant struct {
  //MyColor string
  MyColor string `json:"color"`
}

func (p *Plant) Color() string {
  return p.MyColor
}



// ===== ANIMAL ======
type Animal struct {
  //MyColor string
  MyColor string `json:"color"`
}

func (a *Animal) Color() string {
  return a.MyColor
}

// override the json marshal to specify what you want.
func (a *Animal) MarshalJSON() ([]byte, error) {
	fmt.Println("> animal marshal json called")
  m := make(map[string]string)
  m["type"] = "animal"
  m["color"] = a.MyColor

  return json.Marshal(m)
}


func plants01() {
  // serialize
  p := Plant{MyColor: "green"}
  byteSlice, _ := json.Marshal(p)

  fmt.Println("> plant01 > marshal >", string(byteSlice))

  // deserialize
  newP := Plant{}
  json.Unmarshal(byteSlice, &newP)
  fmt.Println("> plant01 > unmarshal >", newP.MyColor)
}

func animal01() {
  // serialize
  a := Animal{MyColor: "red"}
  byteSlice, _ := json.Marshal(&a)

  fmt.Println("> animal01 > marshal >", string(byteSlice))

  // deserialize
  newA := Animal{}
  json.Unmarshal(byteSlice, &newA)
  fmt.Println("> animal01 > unmarshal color >", newA)
  //fmt.Println("> animal01 > unmarshal type >", newP.Type)
}


func main() {
  plants01()
  animal01()
}



/*
http://gregtrowbridge.com/golang-json-serialization-with-interfaces/

Capitializaton matter on expoting structs.
Capitals matter on json marsahlling automatically

The big thing is that the guy wanted a type in JSON?
And then create a custom marshal and unmarshal func

*/