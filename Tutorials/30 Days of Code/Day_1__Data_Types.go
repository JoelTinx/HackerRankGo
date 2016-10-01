package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var i uint32 = 4
    var d float32 = 4.0
    var s string = "HackerRank "

	var i2 uint32
	var d2 float32
	var s2 string
	
	fmt.Scanf("%d", &i2)
	fmt.Scanf("%f", &d2)
	scanner := bufio.NewReader(os.Stdin)
	s2, _ = scanner.ReadString('\n')
	
	fmt.Println(i + i2)
	fmt.Printf("%0.1f\n", (d + d2))
	fmt.Println(s + s2)
}