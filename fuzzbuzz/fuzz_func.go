package main

import (
    "strconv"
    "fmt"
)
func fizzbuzzjalloh(val int) string {
	if val % 15 ==0{
	return "fizzbuzz"
	} else if val % 5 == 0 {
	return "buzz"
	}else if val % 3 ==0 {
	return "fizz"	
}
return strconv.Itoa(val)
}
func fizzbuzz(num int) []string {
	result := []string{}
	for i := 1; i <= num; i++ {
	result = append(result, fizzbuzzjalloh(i))
	}
return result
}
func main() {
	fmt.Println(fizzbuzz(30))
}
