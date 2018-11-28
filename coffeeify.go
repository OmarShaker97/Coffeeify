package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

func main() {
	w, err := owm.NewCurrent("C", "EN", "9590c142477f0f4ab7b35ec14cf9a446") // celsius (imperial) with English output
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter City: ")
	city, _ := reader.ReadString('\n')
	//fmt.Println(city)

	w.CurrentByName(city)
	fmt.Print("Weather: ")
	fmt.Print(w.Main.Temp)
	fmt.Println(" Celsius")

	if w.Main.Temp >= 25 {
		fmt.Println("Recommended: Cold Drink") // Print out a cold drink from our database Ex: (SELECT * FROM COFFEE C WHERE C.weather = "COLD")
	}

	if w.Main.Temp < 25 {
		fmt.Println("Recommended: Hot Drink") // Print out a hot drink from our database Ex: (SELECT * FROM COFFEE C WHERE C.weather = "HOT")
	}
}
