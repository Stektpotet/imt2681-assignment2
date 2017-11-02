package database

//
// import (
// 	"fmt"
//
// 	"github.com/stektpotet/imt2681-assignment2/fixer"
// )
//
// //CurrencyDB - local storage of stuff
// type CurrencyDB struct {
// 	records map[string]map[string]interface{}
// }
//
// func (db *CurrencyDB) Init() {
// 	db.records = make(map[string]map[string]interface{})
// }
//
// func (db *CurrencyDB) Add(collection string, record interface{}) (err error) {
// 	if db.records == nil {
// 		err = fmt.Errorf("Could not add to records because records is nil")
// 	}
// 	db.records[collection][record.Date] = record
// 	return
// }
//
// func (db *CurrencyDB) Count() int {
// 	return len(db.records)
// }
// func (db *CurrencyDB) Get(key string) (value fixer.Currency, ok bool) {
// 	value, ok = db.records[key]
// 	return
// }
// func (db *CurrencyDB) GetAll() (currencies []fixer.Currency) {
// 	currencies = make([]fixer.Currency, 0, db.Count())
// 	for _, payload := range db.records {
// 		currencies = append(currencies, payload)
// 	}
// 	return
// }
