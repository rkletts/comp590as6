package main

import (
  "fmt"
  "math/big"
  "time"
)

func power_h(n, p int64, res *big.Int) {
  if p == 0 {
    fmt.Printf("Result: %v\n", res)
    return
  }
  nextRes := new(big.Int).Mul(res, bigNewInt(n))
  go power_h(n, p-1, nextRes)
}

func power(n, p int64) {
  go power_h(n, p, big.NewInt(1))
}

func main() {
  n := int64(2)
  p := int64(3)
  power(n, p)
  time.Sleep(3 * time.Second)
}
