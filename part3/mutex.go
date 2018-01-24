package main

import "fmt"
import "time"
import "sync/atomic"
import "sync"
import "math/rand"

func main(){

  var state = make(map[int]int)
  var mutex = &sync.Mutex{}

  var readOps uint64 = 0
  var writeOps uint64 = 0

  for r:=0 ; r <100; r++{
    go func(){
      total := 0
      for {
        key := rand.Intn(5)
        fmt.Println("key",key)
        mutex.Lock()
        total+= state[key]
        mutex.Unlock()
        atomic.AddUint64(&readOps,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }


  for r:=0 ; r <100; r++{
    go func(){
      for {
        key := rand.Intn(5)
        value := rand.Intn(100)
        fmt.Println("key",key)
        mutex.Lock()
        state[key]=value
        mutex.Unlock()
        atomic.AddUint64(&writeOps,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  time.Sleep(time.Second)

  readOpsFinal := atomic.LoadUint64(&readOps)
  writeOpsFinal := atomic.LoadUint64(&writeOps)
  fmt.Println(readOpsFinal,writeOpsFinal)
  fmt.Println(state)
}
