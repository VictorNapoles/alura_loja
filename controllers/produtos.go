package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/alura_loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarToddosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, _ := strconv.ParseFloat(preco, 64)
		quantidadeInt, _ := strconv.Atoi(quantidade)

		models.InserirProdutos(nome, descricao, precoFloat, quantidadeInt)

	}

	http.Redirect(w, r, "/", 301)

}
