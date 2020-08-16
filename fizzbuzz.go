package main

import (
	"fmt"
)

func fizzBuzz() {
    for n := 1; n <= 100; n++ {
        if n%3 == 0 && n%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if n%3 == 0 {
            fmt.Println("Fizz")
        } else if n%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(n)
        }
    }
}
