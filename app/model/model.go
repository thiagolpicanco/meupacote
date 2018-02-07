package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Usuario struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nome    string        `json: "nome" bson:"nome"`
	CPF     string        `json: "cpf" bson: "cpf"`
	Pacotes []Pacote      `json: "pacotes,omitempty" bson: "pacotes"`
}

type Pacote struct {
	Nome    string           `json: "nome,omitempty" bson:"nome"`
	Codigo  string           `json: "codigo,omitempty" bson: "codigo"`
	Eventos []EventoCorreios `json: "eventos,omitempty" bson: "eventos"`
}

type PesquisaCorreios struct {
	Pesquisa   string            `json:"pesquisa"`
	Quantidade string            `json:"quantidade"`
	Resultado  string            `json:"resultado"`
	Objeto     []ObjectoCorreios `json:"objeto"`
}

type ObjectoCorreios struct {
	Numero       string           `json:"numero"`
	Sigla        string           `json:"sigla"`
	Nome         string           `json:"nome"`
	Categoria    string           `json:"categoria"`
	CepDestino   string           `json:"cepDestino"`
	DataPostagem string           `json:"dataPostagem"`
	Eventos      []EventoCorreios `json:"evento"`
}

type EventoCorreios struct {
	Tipo           string            `json:"tipo" bson: "tipo"`
	Status         string            `json:"status" bson: "status"`
	DataOcorrencia string            `json:"data" bson: "data"`
	HoraOcorrencia string            `json:"hora" bson: "hora"`
	Criacao        string            `json:"criacao" bson: "criacao"`
	Descricao      string            `json:"descricao" bson: "descricao"`
	Recebedor      RecebedorCorreios `json:"recebedor,omitempty" bson: "recebedor"`
	Unidade        UnidadeCorreios   `json:"unidade" bson: "unidade"`
}

type UnidadeCorreios struct {
	Local    string           `json:"local" bson: "local"`
	Cidade   string           `json:"cidade" bson: "cidade"`
	UF       string           `json:"uf" bson: "uf"`
	Tipo     string           `json:"tipounidade" bson: "tipo"`
	Endereco EnderecoCorreios `json:"endereco" bson: "endereco"`
}

type RecebedorCorreios struct {
	Nome      string `json:"nome" bson: "nome"`
	Documento string `json:"documento" bson: "documento"`
	Obs       string `json:"comentario" bson: "obs"`
}

type EnderecoCorreios struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Bairro      string `json:"bairro"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
