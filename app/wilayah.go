package app

import (
	"net/http"
	"strconv"

	"../helper"
)

type DataWilayahDetail struct {
	Provinsi_id    int    `json:"provinsi_id"`
	Provinsi_nama  string `json:"provinsi_nama"`
	Kabupaten_id   int    `json:"kabupaten_id,omitempty"`
	Kabupaten_nama string `json:"kabupaten_nama,omitempty"`
	Kecamatan_id   int    `json:"kecamatan_id,omitempty"`
	Kecamatan_nama string `json:"kecamatan_nama,omitempty"`
	Kelurahan_id   int    `json:"kelurahan_id,omitempty"`
	Kelurahan_nama string `json:"kelurahan_nama,omitempty"`
}

type DataWilayah struct {
	Id       int               `json:"id"`
	Kategori string            `json:"kategori"`
	Nama     string            `json:"nama"`
	Rincian  DataWilayahDetail `json:"rincian,omitempty"`
}

func Wilayah(w http.ResponseWriter, r *http.Request) {
	db, err := helper.ConnectDatabase()
	if err != nil {
		helper.HandleError(w, err)
		return
	}
	defer db.Close()

	urlQuery := r.URL.Query()
	kategori := urlQuery.Get("kategori")
	wilayahId, _ := strconv.Atoi(urlQuery.Get("id"))

	selectQuery := ""
	if kategori == "provinsi" {
		selectQuery = "SELECT id AS provinsi_id, nama as provinsi_nama from provinsi where id = ? "
	} else if kategori == "kabupaten" {
		selectQuery =
			`SELECT prov.id AS provinsi_id, prov.nama AS provinsi_nama,
				kab.id AS kabupaten_id, kab.nama AS kabupaten_nama
			FROM provinsi prov
			INNER JOIN kabupaten kab
				ON prov.id = kab.provinsi_id
			WHERE kab.id = ?`
	} else if kategori == "kecamatan" {
		selectQuery =
			`SELECT prov.id AS provinsi_id, prov.nama AS provinsi_nama,
				kab.id AS kabupaten_id, kab.nama AS kabupaten_nama,
				kec.id AS kecamatan_id, kec.nama AS kecamatan_nama
			FROM provinsi prov
			INNER JOIN kabupaten kab
				ON prov.id = kab.provinsi_id
			INNER JOIN kecamatan kec
				ON kab.id = kec.kabupaten_id
			WHERE kec.id = ?`
	} else if kategori == "kelurahan" {
		selectQuery =
			`SELECT prov.id AS provinsi_id, prov.nama AS provinsi_nama,
				kab.id AS kabupaten_id, kab.nama AS kabupaten_nama,
				kec.id AS kecamatan_id, kec.nama AS kecamatan_nama,
				kel.id AS kelurahan_id, kel.nama AS kelurahan_nama
			FROM provinsi prov
			INNER JOIN kabupaten kab
				ON prov.id = kab.provinsi_id
			INNER JOIN kecamatan kec
				ON kab.id = kec.kabupaten_id
			INNER JOIN kelurahan kel
				ON kec.id = kel.kecamatan_id
			WHERE kel.id = ?
			LIMIT 1`
	} else {
		//jika wilayah tidak tersedia return 501
		var emptyData interface{}
		helper.HandleResponse(w, 501, emptyData)
		return
	}

	dataWilayahDetail := DataWilayahDetail{}
	err = db.Get(&dataWilayahDetail, selectQuery, wilayahId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			//jika error karena tidak ada data
			var emptyData interface{}
			helper.HandleResponse(w, 501, emptyData)
			return
		}
		helper.HandleError(w, err)
		return
	}

	dataWilayah := DataWilayah{
		Id:       wilayahId,
		Kategori: kategori,
		Rincian:  dataWilayahDetail,
	}

	switch kategori {
	case "provinsi":
		dataWilayah.Nama = dataWilayahDetail.Provinsi_nama
	case "kabupaten":
		dataWilayah.Nama = dataWilayahDetail.Kabupaten_nama
	case "kecamatan":
		dataWilayah.Nama = dataWilayahDetail.Kecamatan_nama
	case "kelurahan":
		dataWilayah.Nama = dataWilayahDetail.Kelurahan_nama
	}

	helper.HandleResponse(w, 200, dataWilayah)
}
