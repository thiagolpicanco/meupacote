package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/meupacote/app/handlers"
	"github.com/meupacote/config"
	"gopkg.in/mgo.v2"
)

var sessaoBanco *mgo.Session
var router *mux.Router

//Init abre conexao com o mongo e inicializa e seta rotas
func Init(Config *config.Config) {
	var err error
	urlBanco := fmt.Sprintf("mongodb://%s:%s@%s", Config.DB.Username, Config.DB.Password, Config.DB.Host)
	fmt.Println("URL BANCO ", urlBanco)

	sessaoBanco, err = mgo.Dial(urlBanco)

	if nil != err {
		log.Fatal("Erro ao abrir conexao com o mongo ", err)
	}

	router = mux.NewRouter()
	setRouters()

}

func setRouters() {
	router.HandleFunc("/usuario", buscaUsuarios).Methods("GET")
	router.HandleFunc("/usuario", criarUsuario).Methods("POST")
	router.HandleFunc("/hello", helloWorld)
	log.Println("Rotas setadas")
}

func criarUsuario(w http.ResponseWriter, r *http.Request) {
	handlers.CriarUsuario(sessaoBanco, w, r)
}

func buscaUsuarios(w http.ResponseWriter, r *http.Request) {
	handlers.BuscaUsuarios(sessaoBanco, w, r)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	w.Write([]byte(msg))
}

func Run(host string) {
	log.Fatal(http.ListenAndServe(host, router))
}
