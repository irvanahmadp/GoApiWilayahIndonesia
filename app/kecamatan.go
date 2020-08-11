package app

import (
	"../helper"
	"net/http"
)

type DataKecamatan struct {
	Id           int    `json:"id"`
	Kabupaten_id int    `json:"kabupaten_id"`
	Nama         string `json:"nama"`
}

func Kecamatan(w http.ResponseWriter, r *http.Request) {
	db, err := helper.ConnectDatabase()
	if err != nil {
		helper.HandleError(w, err)
		return
	}
	defer db.Close()

	urlQuery := r.URL.Query()
	kabupaten_id := urlQuery.Get("kabupaten_id")

	dataKecamatan := []DataKecamatan{}
	selectQuery := "SELECT id, kabupaten_id, nama FROM kecamatan where kabupaten_id = ?"
	err = db.Select(&dataKecamatan, selectQuery, kabupaten_id)
	if err != nil {
		helper.HandleError(w, err)
		return
	}

	if len(dataKecamatan) == 0 {
		var emptyData interface{}
		helper.HandleResponse(w, 201, emptyData)
		return
	}
	helper.HandleResponse(w, 200, dataKecamatan)
}