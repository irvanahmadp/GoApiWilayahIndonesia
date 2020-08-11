package app

import (
	"../helper"
	"net/http"
)

type DataKabupaten struct {
	Id          int    `json:"id"`
	Provinsi_id int    `json:"provinsi_id"`
	Nama        string `json:"nama"`
}

func Kabupaten(w http.ResponseWriter, r *http.Request) {
	db, err := helper.ConnectDatabase()
	if err != nil {
		helper.HandleError(w, err)
		return
	}
	defer db.Close()

	urlQuery := r.URL.Query()
	provinsi_id := urlQuery.Get("provinsi_id")

	dataKabupaten := []DataKabupaten{}
	selectQuery := "SELECT id, provinsi_id, nama FROM kabupaten where provinsi_id = ? "
	err = db.Select(&dataKabupaten, selectQuery, provinsi_id)
	if err != nil {
		helper.HandleError(w, err)
		return
	}

	if len(dataKabupaten) == 0 {
		var emptyData interface{}
		helper.HandleResponse(w, 501, emptyData)
		return
	}
	helper.HandleResponse(w, 200, dataKabupaten)
}