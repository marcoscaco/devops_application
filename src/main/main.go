# Aplicação mega simples somente para demostrar os conceitos de devOps
# Constste em uma api qual irá consumir uma instancia REDIS para gerar respostas
# Usando Golang devido a sua performance exemplar e simplicidade na construção de um WebServer simples

package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	fmt.Println("INICIANDO SERVIDOR")

	handler.HandleFunc("/api/hello", Hello)
	handler.HandleFunc("/api/status", Helth_check)
	fmt.Println("INICIADO")
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}

func Helth_check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Check OK!`)
}
