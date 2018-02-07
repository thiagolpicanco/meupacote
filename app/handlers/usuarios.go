package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"

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
	id := bson.NewObjectId()
	usuario.ID = id
	err2 := c.Insert(usuario)
	if nil != err2 {

		if strings.Contains(err2.Error(), "duplicate") {
			log.Println(err2)
			respondeErro(w, http.StatusConflict, "Cpf ja existe na base")
		} else {
			log.Println(err2)
			respondeErro(w, http.StatusInternalServerError, err2.Error())
		}
	} else {
		stringID := id.Hex()
		log.Println("Usuario cadastrado ", stringID)
		BuscaUsuario(stringID, sessao, w, r)
	}

}

//BuscaUsuario retorna o usuario pelo ID
func BuscaUsuario(id string, sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {

	copiaSessao := sessao.Copy()
	defer copiaSessao.Close()

	c := copiaSessao.DB("meupacote").C("usuario")
	var resultado model.Usuario
	queryID := bson.ObjectIdHex(id)
	err := c.FindId(queryID).One(&resultado)

	if err != nil {
		respondeErro(w, http.StatusInternalServerError, err.Error())
	}

	respondeJSON(w, http.StatusOK, resultado)
}

//BuscaUsuarios retorna todos os usuarios cadastros na base
func BuscaUsuarios(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {

	copiaSessao := sessao.Copy()
	defer copiaSessao.Close()

	c := copiaSessao.DB("meupacote").C("usuario")
	var results []model.Usuario

	err := c.Find(nil).All(&results)

	if err != nil {
		respondeErro(w, http.StatusInternalServerError, err.Error())
	}

	respondeJSON(w, http.StatusOK, results)
}

//InserePacote adiciona ao usuario um novo pacote
func InserePacote(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pacote model.Pacote

	corpo, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(corpo, &pacote)

	_, encomenda := BuscaEncomenda(pacote.Codigo)
	pesquisa := model.PesquisaCorreios{}
	json.Unmarshal([]byte(encomenda), &pesquisa)
	objetoCorreios := pesquisa.Objeto

	pacote.Eventos = objetoCorreios[0].Eventos
	copiaSessao := sessao.Copy()
	defer copiaSessao.Close()

	c := copiaSessao.DB("meupacote").C("usuario")
	queryID := bson.ObjectIdHex(id)
	queryUpdate := bson.M{"$push": bson.M{"pacotes": pacote}}

	err := c.UpdateId(queryID, queryUpdate)

	if err != nil {
		respondeErro(w, http.StatusInternalServerError, err.Error())
	} else {
		respondeJSON(w, http.StatusOK, pacote)

	}
}

//BuscaPacotes lista todos os pacotes referentes ao usuario
func BuscaPacotes(sessao *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pacotes []model.Pacote

	copiaSessao := sessao.Copy()
	defer copiaSessao.Close()

	var usuario model.Usuario
	c := copiaSessao.DB("meupacote").C("usuario")
	queryID := bson.ObjectIdHex(id)

	err := c.FindId(queryID).One(&usuario)

	if err != nil {
		respondeErro(w, http.StatusInternalServerError, err.Error())
	} else {
		pacotes = usuario.Pacotes
		respondeJSON(w, http.StatusOK, pacotes)

	}
}

//BuscaEncomenda vai na api dos correios buscar dados sobre o pacote
func BuscaEncomenda(_code string) (string, string) {
	apiUrl := "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	data := "<rastroObjeto><usuario>MobileXect</usuario><senha>DRW0#9F$@0</senha><tipo>L</tipo><resultado>T</resultado><objetos>" + _code + "</objetos><lingua>101</lingua><token>QTXFMvu_Z-6XYezP3VbDsKBgSeljSqIysM9x</token></rastroObjeto>"
	client := &http.Client{}
	r, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data))
	if err != nil {
		return "-1 ERROR", "cannot craft request!"
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/xml")
	r.Header.Add("User-Agent", "Dalvik/1.6.0 (Linux; U; Android 4.2.1; LG-P875h Build/JZO34L)")
	resp, err := client.Do(r)
	if err != nil {
		return "-2 ERROR", "cannot send request!"
	}
	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "-3 ERROR", "cannot decode body!"
	}
	return resp.Status, string(_body)
}
