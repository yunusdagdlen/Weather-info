package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"main/application/helpers"
	"main/application/models"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		err := helpers.InsertEmployee(name, city, "")
		if err != nil {
			http.Redirect(w, r, "/401", 301)
		}
	} else {
		http.Redirect(w, r, "/401", 301)
	}
	http.Redirect(w, r, "/", 301)
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	defer db.Close()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	employee := models.Employee{}
	employeeSlice := []models.Employee{}

	for selDB.Next() {
		var id int
		var name, city, weather string
		err = selDB.Scan(&id, &name, &city, &weather)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.City = city
		employee.Weather = weather
		employeeSlice = append(employeeSlice, employee)
	}
	t, _ := template.ParseFiles("templates/index.html")
	_ = t.Execute(w, employeeSlice)

}

func Update(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		_, err := db.Query("UPDATE Employee SET name=$1, city=$2 Where id=$3", name, city, id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("UPDATE: Name:" + name + "| City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := models.Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl := helpers.ParseTemplate()
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	emp := r.URL.Query().Get("id")
	_, err := db.Query("DELETE FROM Employee WHERE id=$1 ", emp)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	nId := r.URL.Query().Get("id")
	selDb, err := db.Query("SELECT * FROM Employee WHERE id=$1", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := models.Employee{}

	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl := helpers.ParseTemplate()

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl := helpers.ParseTemplate()
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Weather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	response, _ := helpers.WeatherApiQuery("1.1.1.1", city)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Current.WeatherDescriptions[0])
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/401.html")
	_ = t.Execute(w, "big error")

}
