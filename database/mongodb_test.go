package database

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestMongoDB_CreateLocalSession(t *testing.T) {

}

func TestMongoDB_CreateSession(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()
	// mock := NewMockDBStorage(mockCtrl)

}

func TestMongoDB_Init(t *testing.T) {

}

func TestMongoDB_ensureIndex(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		s *mgo.Session
		c string
		i mgo.Index
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			db.ensureIndex(tt.args.s, tt.args.c, tt.args.i)
		})
	}
}

func TestMongoDB_Add(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		collection string
		record     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			if err := db.Add(tt.args.collection, tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDB_Count(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		collection string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			if got := db.Count(tt.args.collection); got != tt.want {
				t.Errorf("MongoDB.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDB_Get(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		collection string
		query      bson.M
		data       interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			if gotOk := db.Get(tt.args.collection, tt.args.query, tt.args.data); gotOk != tt.wantOk {
				t.Errorf("MongoDB.Get() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMongoDB_Delete(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		collection string
		query      bson.M
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			if gotOk := db.Delete(tt.args.collection, tt.args.query); gotOk != tt.wantOk {
				t.Errorf("MongoDB.Delete() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMongoDB_GetAll(t *testing.T) {
	type fields struct {
		HostURLs  []string
		AdminUser string
		AdminPass string
		Name      string
	}
	type args struct {
		collection string
		data       interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &MongoDB{
				HostURLs:  tt.fields.HostURLs,
				AdminUser: tt.fields.AdminUser,
				AdminPass: tt.fields.AdminPass,
				Name:      tt.fields.Name,
			}
			db.GetAll(tt.args.collection, tt.args.data)
		})
	}
}
