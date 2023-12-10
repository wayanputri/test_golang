# Aplikasi Catatan (Notes)

## Configurasi
- buatlah file baru dengan nama local.env lalu didalamnya diberikan beberapa credensial untuk connect ke database,
	di aplikasi ini menggunakan mongodb

	```
	export URL=secrect
	export DBNAME=secrect
	export DBCOLL=secrect
	```
- berikut penjelasan mengenai kredensial di atas
	url merupakan link yang dapat menghubungkan aplikasi dengan server, bila anda menggunakan localhost maka bisa menggunakan 
	```
	localhost:27017
	```

	DBNAME merupakan nama database 
	DBCOLL merupakan nama collection

	tetapi di folder ini saya telah menambahkan local.env web, bila ingin langsung mencoba aplikasi tanpa configurasi terlebih dahulu

## Documentasi dan test
- untuk documentasi digunakan swagger, dengan library swago

### mengakses swagger
- jalankan aplikasi dengan mengetik di terminal
	```
	go run main.go
	```
- lalu kita ke browser dengan mengakses
	```
	http://localhost:8080/swagger/index.html
	```
- untuk panduan lengkapnya dapat mengakses 
	* [Swago documentasi](https://github.com/swaggo/swag)

## Menjalankan Testing
- aplikasi ini juga disediakan testing di layer service, untuk menjalankan testing dapat mengetik 
	```
	 go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
	```
## Endpoint Notes
### Membuat catatan baru
```
POST /notes
```
### Mengambil semua catatan
```
GET /notes
```
### Mengambil satu catatan dengan id tertentu
```
GET /notes/:id
```
### Memperbarui catatan dengan id tertentu
```
PUT /notes/:id
```
### Menghapus catatan dengan id tertentu
```
DELETE /notes/:id
```
