package handler

import "net/http"

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {}
