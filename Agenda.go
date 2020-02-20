// API Agenda
// Ivo Dias
// Referencia: https://medium.com/@rafaelacioly/construindo-uma-api-restful-com-go-d6007e4faff6

// Importa o pacote principal
package main

// Importa os pacotes necessarios
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Cria um objeto do tipo Pessoa
type Pessoa struct {
	ID        string    `json:"id,omitempty"`
	Nome      string    `json:"Nome,omitempty"`
	Sobrenome string    `json:"Sobrenome,omitempty"`
	Endereco  *Endereco `json:"Endereco,omitempty"`
}
type Endereco struct {
	Cidade string `json:"Cidade,omitempty"`
	Estado string `json:"Estado,omitempty"`
}

var pessoas []Pessoa

// Mostra todos os contatos da variável pessoas
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

// Mostra apenas um contato
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

// Cria um novo contato
func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var Pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&Pessoa)
	Pessoa.ID = params["id"]
	pessoas = append(pessoas, Pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

// Deleta um contato
func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pessoas {
		if item.ID == params["id"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

// Função principal para executar a api
func main() {
	router := mux.NewRouter()
	pessoas = append(pessoas, Pessoa{ID: "1", Nome: "Joao", Sobrenome: "Silva", Endereco: &Endereco{Cidade: "Sao Paulo", Estado: "SP"}})
	pessoas = append(pessoas, Pessoa{ID: "2", Nome: "Rafael", Sobrenome: "Mafra", Endereco: &Endereco{Cidade: "Brasilia", Estado: "DF"}})
	router.HandleFunc("/contato", GetPessoas).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPessoa).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePessoa).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePessoa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
