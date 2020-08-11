package app

import (
	"../helper"
	"net/http"
)

type DataKelurahan struct {
	Id           int    `json:"id"`
	Kecamatan_id int    `json:"kecamatan_id"`
	Nama         string `json:"nama"`
}

func Kelurahan(w http.ResponseWriter, r *http.Request) {
	db, err := helper.ConnectDatabase()
	if err != nil {
		helper.HandleError(w, err)
		return
	}
	defer db.Close()

	urlQuery := r.URL.Query()
	kecamatan_id := urlQuery.Get("kecamatan_id")

	dataKelurahan := []DataKelurahan{}
	selectQuery := "SELECT id, kecamatan_id, nama FROM kelurahan where kecamatan_id = ?"
	err = db.Select(&dataKelurahan, selectQuery, kecamatan_id)
	if err != nil {
		helper.HandleError(w, err)
		return
	}

	if len(dataKelurahan) == 0 {
		var emptyData interface{}
		helper.HandleResponse(w, 501, emptyData)
		return
	}
	helper.HandleResponse(w, 200, dataKelurahan)
}