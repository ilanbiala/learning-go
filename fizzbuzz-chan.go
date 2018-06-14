package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(count int, c chan<- string) {
	for i := 1; i < count; i++ {
		res := strconv.Itoa(i)

		if (i%3 == 0) && (i%5 == 0) {
			res += " FizzBuzz"
		} else if i%3 == 0 {
			res += " Fizz"
		} else if i%5 == 0 {
			res += " Buzz"
		}

		c <- res
	}
	close(c)
}

func main() {
	count := 20
	c := make(chan string)
	go fizzbuzz(count, c)
	for str_res := range c {
		fmt.Println(str_res)
	}
}
