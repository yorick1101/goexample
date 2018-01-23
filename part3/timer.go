package main

import "fmt"
import "time"

func main(){
  time1 := time.NewTimer(time.Second*2)

  <- time1.C
  fmt.Println("Timer1 expired")

  time2 := time.NewTimer(time.Second)
  go func(){

    <-time2.C
    fmt.Println("Timer2 expired")
  }()

  stop := time2.Stop()
  if stop {
    fmt.Println("Timer2 stopped");
  }

  time.Sleep(time.Second*5)
}
