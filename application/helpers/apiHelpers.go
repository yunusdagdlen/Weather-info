package helpers

import (
	"html/template"
	"log"
)

func ParseTemplate() *template.Template {
	var tmpl = template.Must(template.ParseGlob("form/*"))
	return tmpl
}

func InsertEmployee(name string, city string, weather string) error {
	db := DbConn()
	defer db.Close()

	_, err := db.Query("INSERT INTO Employee (name, city, weather) VALUES($1, $2, $3)", name, city, weather)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Name: " + name + " | City: " + city + "inserted to db.")

	return err

}
