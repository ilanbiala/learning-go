package fizzbuzz

import "fmt"

func main() {
	count := 100
	for i := 1; i < count; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
