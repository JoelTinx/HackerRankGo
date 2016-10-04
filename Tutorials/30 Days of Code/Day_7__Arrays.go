package main

import (
		"bufio"
		"fmt"
		"os"
		"strconv"
		"strings"
	)
	
func main() {
  var N int
  fmt.Scan(&N)
  
  if N >= 1 && N <= 1000 {
  	reader := bufio.NewReader(os.Stdin)
  	tmp, _ := reader.ReadString('\n')
  	A := strings.Split(strings.Replace(tmp, "\n", "", -1), " ")
  	cad := ""
  	for i := N-1; i >= 0; i-- {
  		if val, _ := strconv.Atoi(A[i]); val >= 1 && val <= 1000 {
	  		if i == 0 {
	  			cad += A[i]
	  		} else {
	  			cad += A[i] + " "
	  		}	
  		}
  	}
  	fmt.Println(cad)
  }
}