package handler

import "net/http"

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create order"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List orders"))
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get order by id"))
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update order by id"))
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete order by id"))
}
