package main

import (
	"log"
	"net/http"
)

// HTTP --> é um protocolo de comunicação, é a base da comunicação web e funciona no modelo
// "CLIENTE X SERVIDOR", isso quer dizer que o HTTP funciona num modelo onde um cliente faz
// uma requisição para um servidor, esse servidor processa essa requisição e devolve uma
// resposta pra ela, o cliente faz a requisição (request) e o servidor processa a
// requisição e envia uma resposta (response), essa comunicação acontece atraves de mensagens
// então a requisição na verdade é uma mensagem pro servidor que ele vai interpretar e
// devolver uma outra mensagem para o cliente.

// Rotas --> são uma maneira de você conseguir identificar que tipo de mensagem ta sendo
// enviada e a partir disso identificar que tipo de processamento o servidor vai ter que
// fazer encima dessa mensagem, existem vários tipos de mensagens dentro dentro de uma
// aplicação, mas todas elas vão se comunicar com o servidor usando o mesmo protocolo HTTP,
// o protocolo não muda, o que muda é a requisição e a resposta. A rota serve para que se
// possa identificar que ação deve ser feita e que tipo de mensagem estamos mandando.

// URI --> Identificador do recurso, é uma forma de falar pro servidor do que você ta falando,
// Método --> É usado para que você diga o que você quer fazer com o esse recurso que você
// identificou (URI). Os métodos são: GET, POST, PUT e DELETE.

// GET --> geralmente usado para buscar dados em um recurso
// POST --> geralmente usado para cadastrar dados em um recurso
// PUT --> geralmente usado para atualizar dados em um recurso
// DELETE --> geralmente usado para deletar dados em um recurso

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
