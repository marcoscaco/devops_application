// # Aplicação mega simples somente para demostrar os conceitos de devOps
// # Constste em uma api qual irá consumir uma instancia REDIS para gerar respostas
// # Usando Golang devido a sua performance exemplar e simplicidade na construção de um WebServer simples

package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	fmt.Println("INICIANDO SERVIDOR")

	// Rotas basicas de chack
	handler.HandleFunc("/api/status", Helth_check)

	// Rotas da aplicação CRUD
	handler.HandleFunc("/api/get", get_db)
	handler.HandleFunc("/api/put", insert_db)
	handler.HandleFunc("/api/delete", remove_db)

	http.ListenAndServe("0.0.0.0:8080", handler)
	// Listen to all request to port 8080.
}

func Helth_check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Check OK!`)
}

func insert_db(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `STATUS: INSERIDO!`)
}

func remove_db(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `STATUS: REMOVIDO!`)
}

func get_db(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Check dado`)
}
