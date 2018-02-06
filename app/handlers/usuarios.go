package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/meupacote/app/model"
	"gopkg.in/mgo.v2"
)

var cache []model.Usuario

//CriarUsuario cadastra usuario no banco
func CriarUsuario(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {
	copiaSessao := sessao.Copy()
	defer copiaSessao.Close()
	usuario := model.Usuario{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&usuario); err != nil {
		return
	}
	defer r.Body.Close()
	c := copiaSessao.DB("meupacote").C("usuario")
	err2 := c.Insert(usuario)
	if nil != err2 {

		if strings.Contains(err2.Error(), "duplicate") {
			respondeErro(w, http.StatusConflict, "Cpf ja existe na base")
		} else {
			log.Println(err2)
			respondeErro(w, http.StatusInternalServerError, err2.Error())
		}

	} else {
		respondeJSON(w, http.StatusCreated, usuario)
	}

}

//BuscaUsuarios retorna todos os usuarios cadastros na base
func BuscaUsuarios(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {

	if len(cache) < 1 {
		log.Println("Buscando usuarios do banco e inserindo no cache")
		copiaSessao := sessao.Copy()
		defer copiaSessao.Close()

		c := copiaSessao.DB("meupacote").C("usuario")
		var results []model.Usuario

		err := c.Find(nil).All(&results)
		cache = results

		if err != nil {
			respondeErro(w, http.StatusInternalServerError, err.Error())
		}

	} else {
		log.Println("Usuarios vindos do cache")
	}

	respondeJSON(w, http.StatusOK, cache)
}

func inserePacote(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {

}
