package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Template --> cria templates baseados no nosso arquivo HTML, esse pacote também consegue
// deixar o HTML dinamico, não é um slice, mas contém vários templates dentro dele
var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// atribuimos a variavel 'template' todos os templates criados, ele são identificados
	// pela expressão '*.html', essa expressão faz ele reconhecer a expressão e jogar tudo
	// que puder ser lido dentro do template
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{"João", "joao.pedro@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
