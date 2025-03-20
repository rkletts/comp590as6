package main

import (
  "fmt"
  "math/big"
  "sync"
  "time"
)

var res = big.NewInt(1)
var mu sync.Mutex

func multiply(n int64) {
  mu.Lock()
  defer mu.Unlock()
  res.Mul(res, big.NewInt(n))
}

func power2(n, p int64) {
  for i := int64(0); i < p; i++ { 
    go multiply(n)
  }
}

func main() {
  n := int64(2)
  p := int64(3)
  power2(n, p)
  time.Sleep(3 * time.Second)
  fmt.Printf("Result: %v\n", res)
}
