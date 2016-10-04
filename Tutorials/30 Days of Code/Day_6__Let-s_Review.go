package main

import "fmt"

func main() {
  var (
    T int
    S []string
  )

  fmt.Scan(&T)
      
  if T >= 1 && T <= 10 {
    for i := 0; i < T; i++ {
      var tmp string
      fmt.Scan(&tmp)
      S = append(S, tmp)	    	
    }
    
    for i := 0; i < len(S); i++ {
      if len(S[i]) >= 2 && len(S[i]) <= 10000 {
        var even, odd string = "", ""
        for j := 0; j < len(S[i]); j++ {
          if j % 2 == 0 {
            even += string(S[i][j])
          } else {
            odd += string(S[i][j])
          }
        }
        fmt.Printf("%s %s\n", even, odd)
        even, odd = "", ""
      }	    	
    }      
  }
}