package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

	json.NewEncoder(w).Encode(Livros)
}


func excluirLivro(w http.ResponseWriter, r *http.Request){
	partes := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(partes[2]); if err !=nil{
		//erro
	}
	indiceLivro:=-1
	for indice, livro:= range Livros{
		if livro.Id == id{
			indiceLivro=indice
			break
		}
	}
	if indiceLivro <0{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	
	Livros=append(Livros[0:indiceLivro],Livros[indiceLivro+1:len(Livros)]...)
	w.WriteHeader(http.StatusNoContent)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusCreated)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1
	Livros = append(Livros, novoLivro)
	json.NewEncoder(w).Encode(novoLivro)
}

func rotearLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	partes:= strings.Split(r.URL.Path, "/")

	if len(partes)==2 || len(partes)==3 && partes[2]=="" { 
		if r.Method == "GET" {
			listarLivros(w, r)
		} else if r.Method == "POST" {
			cadastrarLivro(w, r)
		}
	}else if len(partes)==3|| len(partes)==4 && partes[3]==""{
		if r.Method=="GET"{
			buscarLivro(w, r)
		}else if r.Method=="DELETE"{
			excluirLivro(w,r)
		}
		
	}else {
		w.WriteHeader(http.StatusNotFound)
	

}
}

func buscarLivro(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	partes := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(partes[2]); if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, livro:= range Livros{
		if livro.Id == id{
			json.NewEncoder(w).Encode(livro)
		}
	}

	w.WriteHeader(http.StatusNotFound)

}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)
	// /livros/id
	http.HandleFunc("/livros/", rotearLivros)
}
func configurarServidor() {
	configurarRotas()
	fmt.Println("Servidor está rodando na porta 1337")
	http.ListenAndServe(":1337", nil)
}

func main() {
	configurarServidor()
}
