package main

import (
	common "Coffeeify/common"
	"bufio"
	"fmt"
	//"html/template"
	"log"
	"net/http"
	"os"
	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
)

/*func getCity(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("View.html")
	//fmt.Fprintf(w, "Enter your City Please: ")
	var s = "Title: Meh Body: Whatever"
	t.Execute(w, s)

}*/
var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", common.LoginPageHandler) // GET

	router.HandleFunc("/index", common.IndexPageHandler) // GET
	router.HandleFunc("/login", common.LoginHandler).Methods("POST")

	router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")

	http.Handle("/", router)

	http.ListenAndServe(":8000", nil)

	//router.HandleFunc("/weather", getCity).Methods("GET")

	w, err := owm.NewCurrent("C", "EN", "9590c142477f0f4ab7b35ec14cf9a446") // celsius (imperial) with English output
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter City: ")
	city, _ := reader.ReadString('\n')

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
