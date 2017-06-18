package main

import "fmt"

func main() {
  // the string
  s := "hello world    "
  
  // reverse a string
  d1 := make([]byte, len(s))
  copy(d1, s)
  fmt.Println(string(d1))
  for i := 0 ; i < len(d1)/2 ; i++ {
    fmt.Println(i, len(d1)-i)
    d1[i], d1[len(d1)-i-1] = d1[len(d1)-i-1], d1[i] 
  }
  fmt.Println("> char reversed >", string(d1))

  // reverse a string
  // fmt.Println("\n> word reverse...")
  // d2 := make([]byte, len(s))
  // copy(d2, s)
  // iStr := 0
  // for i := 0 ; i < len(d2) ; i++ {
  //   fmt.Println(">", i, len(d2))
  //   // figure out spaces.
  //   if d2[i] == " "[0] || i + 1 == len(d2) {
  //     fmt.Println("> if", iStr+((i-iStr)/2))
  //     for j := iStr ; j < iStr+((i-iStr)/2) ; j++ {
  //       fmt.Println(">>", j, i-j-1)
  //       d2[j], d2[i-j-1] = d2[i-j-1], d2[j]
  //       fmt.Println(string(d2), j, iStr+((i-iStr)/2))
  //     }
  //     iStr = i+1
  //   }
  // }
  // fmt.Println("> word reversed >", string(d2))

  // reverse another words in str
  d2 := make([]byte, 0)
  tmp := make([]byte, 0)
  for i := 0 ; i<len(s) ; i++ {
    if s[i] != " "[0] {
      tmp = append([]byte{s[i]}, tmp...)
    } else {
      d2 = append(d2, tmp...)
      d2 = append(d2, " "[0])
      tmp = make([]byte, 0)
    }
  }

  // appending on end
  if len(tmp) > 0 {
    d2 = append(d2, tmp...)
  }
  fmt.Println("> reverse words in str [", string(d2), "]")
}