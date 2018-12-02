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
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)

/*func getCity(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("View.html")
	//fmt.Fprintf(w, "Enter your City Please: ")
	var s = "Title: Meh Body: Whatever"
	t.Execute(w, s)

}*/
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "123@tcp(db:3306)"
   dbName := "coffee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}
func CreateTabbles(){
	db, err := sql.Open("mysql", "root:123@tcp(db:3306)/coffee")
	if err != nil{
		log.Fatal(err)
	}
	err=db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()
 
	// _,err = db.Exec("CREATE DATABASE ")
	// if err != nil {
	// 	panic(err)
	// }
 
	_,err = db.Exec("USE coffee")
	if err != nil {
		panic(err)
	}
 
	_,err = db.Exec("CREATE TABLE Users (Id int NOT NULL AUTO_INCREMENT PRIMARY KEY, Username Varchar(32), Password varchar(32) )")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE coffee (Id int NOT NULL AUTO_INCREMENT PRIMARY KEY,Name Varchar(100),Recepie Varchar(100),Weather tinyint(1))")
	if err != nil {
		panic(err)
	}

}
var router = mux.NewRouter()

func main() {
	CreateTabbles()
	router.HandleFunc("/", common.LoginPageHandler) // GET

	router.HandleFunc("/index", common.IndexPageHandler) // GET
	router.HandleFunc("/login", common.LoginHandler).Methods("POST")

	router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")

	router.HandleFunc("/displayAll",common.SelcetAllDrinks)
	router.HandleFunc("/show", common.DisplayRecepie)
	router.HandleFunc("/new", common.New)
	router.HandleFunc("/insert", common.InsertDrinks)

	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)

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
