package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/webvillain/bank2/db"
)

func main() {
	db.ConnectDb()
	http.HandleFunc("/users", userHandler)
	http.ListenAndServe(":8080", nil)
}

/*
{
	"Name":"vijay",
	"Email":"vijay@test.com"
}
*/

//var store db.Database = memdb.NewDatabase()
var store *sql.DB = db.DB

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// TODO: crete user
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		u := &db.User{}
		err = json.Unmarshal(data, u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		u, err = store.CreateUser(r.Context(), u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err = json.Marshal(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	}
	if r.Method == http.MethodGet {
		users, err := store.ListUsers(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	}
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("ID")
		err := store.DeleteUser(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("user is deleted"))
	}
	if r.Method == http.MethodPatch {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		u := &db.User{}
		err = json.Unmarshal(data, u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		us, err := store.User(r.Context(), u.ID)
		if errors.Is(err, db.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		changed := false
		if u.Name != us.Name {
			changed = true
		}
		if u.Email != us.Email {
			changed = true
		}
		if !changed {
			return
		}
		err = store.UpdateUser(r.Context(), u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err = json.Marshal(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(data)
	}
}

// i changed rw (response writer)to w(responsewriter) for my understandings
