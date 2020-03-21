package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

var indexTpl = template.Must(template.ParseFiles("index.html"))

var Index = func(w http.ResponseWriter, r *http.Request) {
	err := indexTpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Ошибка индекса!!!")
	}
}
