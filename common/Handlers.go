package handlers

import (
    helpers "Coffeeify/helpers"
 //   repos "Coffeeify/repos"
    "fmt"
    "github.com/gorilla/securecookie"
    "html/template"
    "net/http"
    _"github.com/go-sql-driver/mysql"
    owm "github.com/briandowns/openweathermap"
	"database/sql"
)


func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password@tcp(localhost:3306)"
    dbName := "coffee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

type drink struct{
	Id int
	Name string
	Recepie string
	Weather bool
}

var tmpl=template.Must(template.ParseGlob("templates/*"))

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

// Handlers

// for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/login.html")
    fmt.Fprintf(response, body)
}

// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
    db := dbConn()
    name := request.FormValue("name")
    pass := request.FormValue("password")

    redirectTarget := "/"
    if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
        // Database check for user data!
        var count = 0

        selDb, err := db.Query("SELECT * FROM Users where Username = ? AND Password = ?", name, pass)

        if(err!=nil){
            panic(err.Error())
        }
        for selDb.Next(){
            count++
            selDb.Scan()
        }

       // _userIsValid := repos.UserIsValid(name, pass)

        if count > 0 {
            SetCookie(name, response)
            redirectTarget = "/displayAll"

        } else {
            redirectTarget = "/register"

        }
    }

    defer db.Close()

    http.Redirect(response, request, redirectTarget, 302)
}

// for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/register.html")
    fmt.Fprintf(response, body)
}

// for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    db :=dbConn()
    r.ParseForm()
    redirectTarget := "/"
    uName := r.FormValue("username")
    email := r.FormValue("email")
    pwd := r.FormValue("password")
    confirmPwd := r.FormValue("confirmPassword")

    _uName, _email, _pwd, _confirmPwd := false, false, false, false
    _uName = !helpers.IsEmpty(uName)
    _email = !helpers.IsEmpty(email)
    _pwd = !helpers.IsEmpty(pwd)
    _confirmPwd = !helpers.IsEmpty(confirmPwd)

    if _uName && _email && _pwd && _confirmPwd {
       // fmt.Fprintln(w, "Username for Register : ", uName)
       // fmt.Fprintln(w, "Email for Register : ", email)
       // fmt.Fprintln(w, "Password for Register : ", pwd)
       // fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
        insForm, err := db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
        
        if err != nil {
            panic(err.Error())
        }
    
        insForm.Exec(uName,pwd)
        redirectTarget = "/displayAll"

        } else {
        // fmt.Fprintln(w, "This fields can not be blank!")
        redirectTarget = "/register"
        }
        
        defer db.Close()

        http.Redirect(w, r, redirectTarget, 302)
    }
   



// for GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
    userName := GetUserName(request)
    if !helpers.IsEmpty(userName) {
        var indexBody, _ = helpers.LoadFile("templates/index.html")
        fmt.Fprintf(response, indexBody, userName)
    } else {
        http.Redirect(response, request, "/", 302)
    }
}

// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
    ClearCookie(response)
    http.Redirect(response, request, "/", 302)
}

// Cookie

func SetCookie(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "name": userName,
    }
    if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
        cookie := &http.Cookie{
            Name:  "cookie",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}

func ClearCookie(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "cookie",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}


///////////////////////////////////////////////////


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
	}
	
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}
func RecommendDrinksPageHandler(response http.ResponseWriter, request *http.Request){
    var body, _ = helpers.LoadFile("templates/recommendDrinks.html")
    fmt.Fprintf(response, body)

}

func RecommendDrinks(response http.ResponseWriter, request *http.Request){

    db := dbConn()
    city := request.FormValue("city")

    //http.Redirect(response, request, redirectTarget, 302)

    w, err := owm.NewCurrent("C", "EN", "9590c142477f0f4ab7b35ec14cf9a446") // celsius (imperial) with English output
	if err != nil {
		fmt.Print(err)
	}
    if (request.Method == "POST"){
	w.CurrentByName(city)
	fmt.Fprintf(response,"Weather: ")
	fmt.Fprintf(response,"%f", w.Main.Temp)
	fmt.Fprintf(response," Celsius \n")

	if w.Main.Temp >= 25 {
        selDB, err := db.Query("SELECT * FROM coffee WHERE Weather=false ORDER BY RAND() LIMIT 1") 
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
        
        fmt.Fprintf(response, "Your recommended drink is %s, and the recepie is %s", Name, Recepie)

		allDrinks =append(allDrinks,currentDrink)
	}
       // fmt.Fprintf(response,"Recommended: Cold Drink")
        tmpl.ExecuteTemplate(response, "recommendDrinks", allDrinks) // Print out a cold drink from our database Ex: (SELECT * FROM COFFEE C WHERE C.weather = "COLD")
	}

	if w.Main.Temp < 25 {

        selDB, err := db.Query("SELECT * FROM coffee WHERE Weather=true ORDER BY RAND() LIMIT 1") 
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
        
        fmt.Fprintf(response, "Your recommended drink is %s, and the recepie is %s", Name, Recepie)

		allDrinks =append(allDrinks,currentDrink)
	}


     //   fmt.Fprintf(response,"Recommended: Hot Drink")
        tmpl.ExecuteTemplate(response, "recommendDrinks", allDrinks) // Print out a hot drink from our database Ex: (SELECT * FROM COFFEE C WHERE C.weather = "HOT")
    }

}
    defer db.Close()

}


