package main

import (
	"bufio"
	"fmt"
	"os"
	)

func main() {
	inputString := "Hello, World!"
    fmt.Println(inputString)
    
    reader := bufio.NewReader(os.Stdin)
    inputString, _ = reader.ReadString('\n')
    fmt.Println(inputString)
}