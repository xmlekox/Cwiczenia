package main

import "fmt"

func main() {
	var imie string
	fmt.Printf("Podaj swoje imie: ")
	_, _ = fmt.Scanf("%s", &imie)
	var il int64
	il = int64(len(imie))
	if imie[il-1] == 'a' {
		fmt.Printf("To jest imię żeńskie!\n")
	} else {
		fmt.Printf("To jest imie męskie!\n")
	}

}
