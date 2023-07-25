package application

import (
	"main/application/controllers"
	"net/http"
)

func InitializeEndpoints() {
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/weather", controllers.Weather)
	http.HandleFunc("/401", controllers.ErrorPage)
}
