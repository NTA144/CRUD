package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Livros []Livro = []Livro{
	{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "José de Alencar",
	}, {
		Id:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	}, {
		Id:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "BEM VINDO")
}
func listarLivros(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Livros)
}
func rotearLivros(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listarLivros(w, r)
	}else if 

}
func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)
	//http.HandleFunc("/livros", cadastrarLivro)
}
func configurarServidor() {
	configurarRotas()
	fmt.Print("Servidor está rodando na porta 1337")
	http.ListenAndServe(":1337", nil)
}

func main() {
	configurarServidor()

}
