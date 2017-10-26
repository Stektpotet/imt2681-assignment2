package database

import (
	"fmt"

	"github.com/stektpotet/imt2681-assignment2/fixer"
)

type CurrencyDB struct {
	records map[string]fixer.Currency
}

func (db *CurrencyDB) Init() {
	db.records = make(map[string]fixer.Currency)
}

func (db *CurrencyDB) Add(record fixer.Currency) error {
	if db.records == nil {
		return fmt.Errorf("Could not add to records because records = nil")
	}
	db.records[record.Date] = record
	return nil
}

func (db *CurrencyDB) Count() uint {
	return 0
}
func (db *CurrencyDB) Get(key string) (value fixer.Currency, ok bool) {
	value, ok = db.records[key]
	return
}
func (db *CurrencyDB) GetAll() []fixer.Currency {
	all := make([]fixer.Currency, 0, db.Count())
	for _, payload := range db.records {
		all = append(all, payload)
	}
	return all
}
