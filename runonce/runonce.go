package runonce

import (
	
	"log"
	
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)
func CreateDB(){

	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/coffee")
	if err != nil{
		log.Fatal(err)
	}
	err=db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()
	 _,err = db.Exec("CREATE DATABASE coffee")
	 if err != nil {
	 	panic(err)
	 }
}

func CreateTables(){
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/coffee")
	if err != nil{
		log.Fatal(err)
	}
	err=db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()
 

 
	_,err = db.Exec("USE coffee_db")
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
func main(){
	//CreateDB()
	CreateTables()
}