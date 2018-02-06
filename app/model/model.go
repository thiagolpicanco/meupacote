package model

import "gopkg.in/mgo.v2/bson"

type Usuario struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Nome string        `json: "nome" bson:"nome"`
	CPF  string        `json: "cpf" bson: "cpf"`
}
