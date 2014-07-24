package main

import "fmt"

func main() {
	// Z dokumentacji: http://golang.org/ref/spec#Numeric_types
	// uint8 the set of all unsigned  8-bit integers (0 to 255)
	var a int16 = 32767 // maksymalna wartosc
	var b int16 = 32767
	var sum int16 = a + b // jaka bedzie wartosc sum?
	fmt.Printf("a=%d, b=%d, a+b=%d\n", a, b, sum)
}
