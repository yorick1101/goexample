package main

import (
  "fmt";
  "strconv"
)

var p = fmt.Println

func main(){

  f,_ := strconv.ParseFloat("1.234", 64)
  p(f)

  i,_ := strconv.ParseInt("1234", 0, 64)
  p(i)

  i2,_ := strconv.ParseInt("0x1cb", 0, 64)
  p(i2)

  u,_ := strconv.ParseUint("789", 0, 64)
  p(u)

  k, _ := strconv.Atoi("135")
  fmt.Println(k)

  k2, e := strconv.Atoi("wat")
  if e != nil {
    panic(e)
  }
  p(k2)
}
