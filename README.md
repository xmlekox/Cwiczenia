Projekt-2-problem
=================
package main

import "fmt"

func main() {

	var distance int
	var road_type int

	fmt.Printf("Prosze podac dystans: ")
	_, _ = fmt.Scanf("%d", &distance)

	fmt.Println("1) Obszar zabudowany (50km/h)")
	fmt.Println("2) Obszar niezabudowany, jednojezdniowa (90km/h)")
	fmt.Println("3) Dwujezdniowa lub ekspresowa jednojezdniowa (100km/h)")
	fmt.Println("4) Ekspresowa dwujezdjiowa (120km/h)")
	fmt.Println("5) Autostrada (140km/h)")

	fmt.Printf("Proszę podać typ drogi:")
	_, _ = fmt.Scanf("%d", &road_type)
	fmt.Printf("Twoja drogra to numer: %d\n", road_type)

	var speed_limit int
	if road_type == 1 {
		speed_limit = 50
	} else if road_type == 2 {
		speed_limit = 90
	} else if road_type == 3 {
		speed_limit = 100
	} else if road_type == 4 {
		speed_limit = 120
	} else if road_type == 5 {
		speed_limit = 140
	}

	hours := distance / speed_limit
	minutes := 60 * float64(distance%speed_limit) / float64(speed_limit)
	fmt.Printf("Hours: %d \n", hours)
	fmt.Printf("Minutes: %d \n", minutes)
}
