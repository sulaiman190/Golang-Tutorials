package main

import "fmt"

func FizzBuzz(x int){
    for i := 1; i <= x; i++ {
        // For multiples of both three and five, print "FizzBuzz".
        if i%15 == 0 {
            fmt.Println("FizzBuzz")
            fmt.Println("---------------------------------------------")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
            fmt.Println("---------------------------------------------")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
            fmt.Println("---------------------------------------------")
        } else {
    
            fmt.Println(i)
        }
    }

}

func main() {
    FizzBuzz(20)



}