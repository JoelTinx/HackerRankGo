package main

import "fmt"

func main() {
  var n int
  fmt.Scanf("%d", &n)
  if n >=2 && n <= 20 {
    for i := 1; i <= 10; i++ {
  	  fmt.Printf("%d x %d = %d\n", n, i, n * i)
    }    
  }
}