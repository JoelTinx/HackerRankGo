package main

import "fmt"

func main() {
  var n uint8
  fmt.Scanf("%d", &n)
  if n >=1 && n <= 100 {
  	if n % 2 != 0 {
	  fmt.Println("Weird")
	} else {
	  if n >= 2 && n <= 5 {
	  	fmt.Println("Not Weird")
	  }  else if n >= 6 && n <= 20 {
	  	fmt.Println("Weird")
	  } else if n > 20 {
	  	fmt.Println("Not Weird")
	  }
	}
  }
}