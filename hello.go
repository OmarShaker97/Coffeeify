package main

import(
	"database/sql"
	"log"
	"html/template"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	http.HandleFunc("/",SelcetAllDrinks)
	http.HandleFunc("/show", DisplayRecepie)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", InsertDrinks)
	http.ListenAndServe(":33060", nil)
}

type drink struct{
	Id int
	Name string
	Recepie string
	Weather bool
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "123@tcp(localhost:3306)"
    dbName := "coffee_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

var tmpl=template.Must(template.ParseGlob("Form/*"))



func SelcetAllDrinks(w http.ResponseWriter, r *http.Request){
	db :=dbConn()
	selDB, err := db.Query("SELECT * FROM coffee ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
	}
	currentDrink := drink{}
	allDrinks := []drink{}

	for selDB.Next() {
        var Id int
		var Name, Recepie string
		var Weather bool
        err = selDB.Scan(&Id,&Name, &Recepie,&Weather)
	
		if err != nil {
            panic(err.Error())
		}
		
		currentDrink.Id=Id
		currentDrink.Name = Name
		currentDrink.Recepie= Recepie
		currentDrink.Weather=Weather
		allDrinks =append(allDrinks,currentDrink)
	}
	
	tmpl.ExecuteTemplate(w, "Index", allDrinks)
	defer db.Close()
}


func DisplayRecepie(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	rowID := r.URL.Query().Get("id")
	selDB,err:=db.Query("SELECT * FROM coffee WHERE Id=?",rowID)
	if err != nil {
        panic(err.Error())
	}
	drink := drink{}

	for selDB.Next(){
		var Id int
		var Name, Recepie string
		var Weather bool
        err = selDB.Scan(&Id,&Name, &Recepie,&Weather)
		if err != nil {
            panic(err.Error())
		}

        drink.Name = Name
		drink.Recepie = Recepie
	}

	tmpl.ExecuteTemplate(w,"Show", drink)
    defer db.Close()
}


func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func InsertDrinks(w http.ResponseWriter, r *http.Request){
	db :=dbConn()
	if (r.Method == "POST"){
		
		Name:= r.FormValue("name")
		Recepie:=r.FormValue("Recepie")
		Weather :=r.FormValue("Weather")
	
	insForm, err := db.Prepare("INSERT INTO coffee (Name ,Recepie ,Weather) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(Name, Recepie, Weather)
	log.Println("INSERT: Drink Name: " + Name + " | Recepie " + Recepie + " | Weather " + Weather) 
	}
	
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}













//Already Created I want to delet this function to avoid further database mapulatuion
func CreateTabbles(){
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/coffee_db")
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

func InsertDrink(){
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/coffee_db")

	Insert, err := db.Query("INSERT INTO coffee(Name ,Recepie ,Weather) Values('Cappucino','1 cup of milk , 2 shots',1)")
	if err != nil {
		panic(err)
	}
	Insert.Close()
}


func RegesiterUser(Username string,Password string){
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/coffee_db")
	Insert, err := db.Query("INSERT INTO Users (Username ,Password) Values('Isra Ragheb','123')")
	if err != nil {
		panic(err)
	}
	Insert.Close()
}

//Already Created I want to delet this function to avoid further database mapulatuion
func DropTables(){
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/coffee_db")

	if err != nil{
		log.Fatal(err)
	}

	_,err = db.Exec("Drop table coffee")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("Drop table Users")
	if err != nil {
		panic(err)
	}
}