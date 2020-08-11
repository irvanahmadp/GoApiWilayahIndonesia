package app

import (
	"../helper"
	"net/http"
)

type DataProvinsi struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

func Provinsi(w http.ResponseWriter, r *http.Request) {
	db, err := helper.ConnectDatabase()
	if err != nil {
		helper.HandleError(w, err)
		return
	}
	defer db.Close()

	dataProvinsi := []DataProvinsi{}
	query := "SELECT id, nama FROM provinsi"
	err = db.Select(&dataProvinsi, query)
	if err != nil {
		helper.HandleError(w, err)
		return
	}

	if len(dataProvinsi) == 0 {
		var emptyData interface{}
		helper.HandleResponse(w, 501, emptyData)
		return
	}
	helper.HandleResponse(w, 200, dataProvinsi)
}