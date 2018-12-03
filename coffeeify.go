package main

import (
	common "Coffeeify/common"
	//"html/template"
	"log"
	"net/http"
	"os"
	"fmt"
	// Shortening the import reference name seems to make it a bit easier
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)


func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp(db:3306)"
	dbName := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}
func CreateTabbles(){
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp(db:3306)/"+os.Getenv("MYSQL_DATABASE"))
	if err != nil{
		log.Fatal(err)
	}
	err=db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()
 
	//  _,err = db.Exec("CREATE DATABASE coffee")
	//  if err != nil {
	//  	panic(err)
	//  }
 
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
	
	fmt.Println(os.Getenv("MYSQL_USER"))
	fmt.Println(os.Getenv("MYSQL_ROOT_PASSWORD"))
	
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
	router.HandleFunc("/recommendDrink",common.RecommendDrinksPageHandler).Methods("GET")
	router.HandleFunc("/recommendDrink", common.RecommendDrinks).Methods("POST")

	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
