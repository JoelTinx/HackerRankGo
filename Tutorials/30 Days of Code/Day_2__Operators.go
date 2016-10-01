package main

import "fmt"

func main() {
  var mealCost   float64 
  var tipPercent int
  var taxPercent int
  
  fmt.Scanf("%f", &mealCost)
  fmt.Scanf("%d", &tipPercent)
  fmt.Scanf("%d", &taxPercent)
  
  tip := mealCost * (float64(tipPercent) / 100)
  tax := mealCost * (float64(taxPercent) / 100)
    
  totalCost := mealCost + tip + tax
  fmt.Printf("The total meal cost is %.0f dollars.", totalCost)
}