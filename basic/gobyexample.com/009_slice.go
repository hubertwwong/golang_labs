package main

import "fmt"
import "strconv"
// needed to the string to int conversion.

func main() {
  s := make([]string, 3)
  fmt.Println("slice intial", s)

  // set the slice items to a vlaue.
  s[0] = "aaaa"
  s[1] = "bbbb"
  s[2] = "ccccc"

  fmt.Println("s", s)
  fmt.Println("s0 ", s[0])

  // you can append after a slice
  sAppend2 := append(s, "3232")
  sAppend := append(sAppend2, "end")
  fmt.Println("sAppend", sAppend)

  // slices can be copied by items
  copied := make([]string, len(s))
  copy(copied, s)
  fmt.Println("copied slice.", copied)

  // probalby the reason slices are they way they are
  s_mid := sAppend[2:3]
  fmt.Println("s_mid ", s_mid)
  s_start := sAppend[2:]
  fmt.Println("s_start ", s_start)
  s_end := sAppend[:2]
  fmt.Println("s_end ", s_end)

  // slice shorthand.
  s_short := []string{"1", "2", "3"}
  fmt.Println("s_shorthand", s_short)

  // 2d slices..
  // can be different unlike
  twoD := make([][]string, 2)
  for i:=0 ; i<len(twoD) ; i++ {
    lenInner := i+3
    // note that this is not using :=
    twoD[i] = make([]string, lenInner)
    for j:=0 ; j<lenInner ; j++ {
      twoD[i][j] =  strconv.Itoa(i) + " " +  strconv.Itoa(j)
    }
  }
  fmt.Println("twodees", twoD)
}
