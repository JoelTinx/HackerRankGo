package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	)

func main() {
  var size int
  fmt.Scanf("%d", &size)
  
  reader := bufio.NewReader(os.Stdin)
  inputNumbers, _ := reader.ReadString('\n')
  numbers, _ := func() ([]int, error) {
  	tmp := strings.Split(strings.Replace(inputNumbers, "\n", "", -1), " ")
  	tmpn := tmp[0:len(tmp)]
  	return convertArrayStringToInt(tmpn)
  }()
  
  
  if len(numbers) >= size {
  	sum := 0
  	for _, num := range numbers {
  		sum += num
  	}
  	fmt.Println(sum)
  } else {
  	fmt.Println("The conditions are met")
  }
}

func convertArrayStringToInt(letters []string) ([]int, error) {
	var err error
	var nums []int
	for _, s := range letters {
		temp, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, temp)
	}
	return nums, err
}