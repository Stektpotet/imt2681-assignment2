package database

import (
	"github.com/stektpotet/imt2681-assignment2/fixer"
	"gopkg.in/mgo.v2/bson"
)

type CurrencyOut struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	*fixer.Currency
}
