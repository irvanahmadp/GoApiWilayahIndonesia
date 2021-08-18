# API Wilayah Indonesia

Rest API wilayah Indonesia dengan golang.

Untuk menjalankan REST API:

1. Import database dari folder db_example.
2. Setting username, password, server database di `helper/database.go`.
3. Jalankan aplikasi dengan perintah `go run main.go` atau build aplikasi `go build main.go` dan jalankan hasil build.

## Menjalankan Aplikasi

```go
go run main.go
```

## Build Aplikasi

```go
go build main.go
```

## Rest API

Contoh request dan response dari REST API ada di bawah.

### Daftar Provinsi

Request

`GET /provinsi/`

Response

```json
{
   "result":200,
   "status":"success",
   "pesan":"",
   "data":[
      {
         "id":11,
         "nama":"ACEH"
      }
   ]
}
```

### Daftar Kabupaten

Request

`GET /kabupaten/?provinsi_id=<provinsi_id>`

Response

```json
{
   "result":200,
   "status":"success",
   "pesan":"",
   "data":[
      {
         "id":1101,
         "provinsi_id":11,
         "nama":"KABUPATEN SIMEULUE"
      }
   ]
}
```

### Daftar Kecamatan

Request

`GET /kecamatan/?kabupaten_id=<kabupaten_id>`

Response

```json
{
   "result":200,
   "status":"success",
   "pesan":"",
   "data":[
      {
         "id":1101010,
         "kabupaten_id":1101,
         "nama":"TEUPAH SELATAN"
      }
   ]
}
```

### Daftar Kelurahan

Request

`GET /kelurahan/?kecamatan_id=<kecamatan_id>`

Response

```json
{
   "result":200,
   "status":"success",
   "pesan":"",
   "data":[
      {
         "id":1101010001,
         "kecamatan_id":1101010,
         "nama":"LATIUNG"
      }
   ]
}
```
