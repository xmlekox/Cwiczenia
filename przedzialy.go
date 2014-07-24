package main

import "fmt"

func main() {
	var liczba int
	fmt.Print("Podaj liczbe: ")
	fmt.Scanf("%d", &liczba)
	fmt.Println("(10, 15): 10 < liczba && liczba < 15 ==", 10 < liczba && liczba < 15)
	fmt.Println("[10, 15): 10 <= liczba && liczba < 15 ==", 10 <= liczba && liczba < 15)
	fmt.Println("[10, 15]: 10 <= liczba && liczba <= 15 == ", 10 <= liczba && liczba <= 15)
	fmt.Println("(10, 15]: 10 < liczba && liczba <= 15 ==", 10 < liczba && liczba <= 15)
}

