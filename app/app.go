package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/thiagolpicanco/meupacote/app/handlers"
	"github.com/thiagolpicanco/meupacote/config"
	"gopkg.in/mgo.v2"
)

var sessaoBanco *mgo.Session
var router *mux.Router

//Init abre conexao com o mongo e inicializa e seta rotas
func Init(Config *config.Config) {
	var err error
	urlBanco := fmt.Sprintf("mongodb://%s:%s@%s", Config.DB.Username, Config.DB.Password, Config.DB.Host)
	sessaoBanco, err = mgo.Dial(urlBanco)

	if nil != err {
		log.Fatal("Erro ao abrir conexao com o mongo ", err)
	}

	router = mux.NewRouter()
	setRouters()

}

func setRouters() {
	Get("/usuario", buscaUsuarios)
	Get("/usuario", buscaUsuarios)
	Get("/usuario/{id}", buscaUsuario)
	Post("/usuario", criarUsuario)
	Put("/usuario/{id}/pacote", adicionaPacote)
	Get("/usuario/{id}/pacote", buscaPacotes)
	Get("/hello", helloWorld)

}

// Get wraper para as reqs GET
func Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("Rota Adicionada: ", path)
	router.HandleFunc(path, f).Methods("GET")
}

// Post  wraper para as reqs POST
func Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("Rota Adicionada: ", path)
	router.HandleFunc(path, f).Methods("POST")
}

//  wraper para as reqs PUT
func Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("Rota Adicionada: ", path)
	router.HandleFunc(path, f).Methods("PUT")
}

// wraper para as reqs DELETE
func Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("Rota Adicionada: ", path)
	router.HandleFunc(path, f).Methods("DELETE")
}

func criarUsuario(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	handlers.CriarUsuario(sessaoBanco, w, r)
}

func buscaUsuarios(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	handlers.BuscaUsuarios(sessaoBanco, w, r)
}

func buscaUsuario(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	vars := mux.Vars(r)
	id := vars["id"]
	handlers.BuscaUsuario(id, sessaoBanco, w, r)
}

func buscaPacotes(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	handlers.BuscaPacotes(sessaoBanco, w, r)
}

func adicionaPacote(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	handlers.InserePacote(sessaoBanco, w, r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	logaRequisciao(r)
	msg := "Hello World"
	w.Write([]byte(msg))
}

func logaRequisciao(r *http.Request) {
	log.Println("Recebendo Requisicao ", r.Method, " na rota:", "{", r.URL, "}", "Ip Requiscao :", r.RemoteAddr)
}

func Run(host string) {
	log.Fatal(http.ListenAndServe(host, router))
}
