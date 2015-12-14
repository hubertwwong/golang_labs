package main

import "fmt"
//import "strconv"

func main() {
  // note that maps are hashes...
  a := make(map[string]int)
  a["foo"] = 1212
  a["bar"] = 12121

  fmt.Println(a)

  b := a["foo"]
  fmt.Println("b gets assigned to a[foo]", b)

  fmt.Println("length", len(a))

  delete(a, "bar")

  fmt.Println("a after deleting foo", a)

  // so 2 things.
  // _ is to ignore the arg on return
  // maps returns 2 things.
  // 1. the value
  // 2. if the key is in the system.
  _,prs := a["baz"]
  fmt.Println("baz was present?", prs)

  // short init
  d := map[string]int {"fiz": 1, "biz": 2}
  fmt.Println("shortcut init was: ", d)
}
