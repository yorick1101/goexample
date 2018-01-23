package main

import "fmt"
import "time"

func main(){
  time1 := time.NewTicker(time.Millisecond*500)


  go func(){
    for t := range time1.C {
      fmt.Println("Ticker fired", t)
    }

  }()

  time.Sleep(time.Millisecond * 1600)
  time1.Stop()
  fmt.Println("Ticker stopped")

}
