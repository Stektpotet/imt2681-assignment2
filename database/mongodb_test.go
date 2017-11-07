package database

import (
	"os"
	"testing"

	"github.com/Stektpotet/imt2681-assignment2/fixer"
	"github.com/golang/mock/gomock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var DB DBStorage

func TestMain(m *testing.M) {

	c := m.Run()

	os.Exit(c)
}

func TestMongoDB_CreateLocalSession(t *testing.T) {

}

func TestMongoDB_CreateSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := NewMockDBStorage(mockCtrl)

	DB = mock

	mock.EXPECT().CreateSession()

	//Invoke
	mock.CreateSession()
}

func TestMongoDB_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockDBStorage(ctrl)
	DB = mock

	mock.EXPECT().Init()

	//Invoke Init
	mock.Init()

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockDBStorage(ctrl)
	DB = mock
	latest, err := fixer.GetLatest("")
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Add("currency", latest)

	//Invoke Init
	mock.Add("currency", latest)
}

func TestMongoDB_Count(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := NewMockDBStorage(mockCtrl)

	DB = mock
	mock.EXPECT().Count("currency")

	//Invoke
	mock.Count("currency")
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
