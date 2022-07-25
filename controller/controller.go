package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/index.html", "template/header.html", "template/body.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}

func Technologies(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/technologies/technologies.html", "template/header.html", "template/technologies/techSkills.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}

func Probien(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/projects/probien/probien.html", "template/projects/probien/probienBody.html", "template/header.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}

func Gespro(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/projects/gespro/gespro.html", "template/header.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}

func Covec(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/projects/covec/covec.html", "template/header.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}

func Pibe(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/projects/pibe/pibe.html", "template/header.html", "template/footer.html")
	if err != nil {
		log.Print(err)
	}
	tpl.Execute(w, nil)
}
