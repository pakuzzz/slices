package main

import (
	"fmt"
)

func addCity(cities []string, city string) []string {
	return append(cities, city)
}

func removeCity(cities []string, city string) []string {
	for i, v := range cities {
		if v == city {
			return append(cities[:i], cities[i+1:]...)
		}
	}
	return cities
}

func containsCity(cities []string, city string) bool {
	for _, v := range cities {
		if v == city {
			return true
		}
	}
	return false
}

func main() {
	cities := []string{"Москва", "Париж", "Лондон"}

	cities = addCity(cities, "Берлин")
	fmt.Println("После добавления:", cities)

	cities = removeCity(cities, "Париж")
	fmt.Println("После удаления:", cities)

	fmt.Println("Есть ли Лондон?", containsCity(cities, "Лондон"))
	fmt.Println("Есть ли Нью-Йорк?", containsCity(cities, "Нью-Йорк"))
}
