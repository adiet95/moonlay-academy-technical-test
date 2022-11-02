
## Moonlay Academy -  Technical Test

Selamat datang di **Moonlay Academy - Technical Test.**

# To Do App
Mini App for list todo and add some sublist basely with Golang, Echo v4 REST API Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description
```bash
### Spesifikasi Bisnis :
- Menampilkan, menambahkan, mengubah dan menghapus data.
- Dapat menambahkan sub list untuk setiap list yang terdaftar. Sub list bisa ditambah, diubah dan dihapus. ( hanya 2 level )

Data Input untuk masing-masing list/sub list :
- Title [required | maximum 100 karakter | alphanumeric]
- Deskripsi [required | maximum 1000 karakter]
- File Upload [optional | hanya menerima file dengan extension .txt dan .pdf]

Buat API Todo List dengan kriteria sebagai berikut :

Menampilkan data list tanpa sub list ( dengan dan tanpa pagination ).
Menampilkan data list beserta dengan sub list nya jika ada.
Menampilkan data sub list by list id.
Menambahkan data list.
Menambahkan data sub list untuk spesifik list.
Mengubah data list/sub list dengan kritera input diatas.
Menghapus data list/sub list.
```

## Notes :
1. In POST & PUT List can attach .pdf / .txt file, use Form Data for it.
2. In POST & PUT **Sublist can't attach** .pdf/.txt, use data raw for it, There are some bug to attach in sublist
3. In this case, I have implemented DI (Dependecies Injection) there are 3 layer on it, Repo for quering to DB, Service to handle bussiness logic before sending to repo, Controller to Request data from body or query and Respons data.

## Several command you must know in this app :
```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
3. go run . migrate -d //for database rollback
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://github.com/adiet95/moonlay-academy-technical-test.git
```
2. Install dependencies
```bash
go mod tidy
```
> Wait a minute, if still error run :

```bash
go mod vendor
```

3. Add Env File

```sh
  USER = Your DB User
  HOST = Your DB Host
  DB   = Your DB Name
  PASS = Your DB Password
  PORT = Your Port
```

4. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration table
```

5. Run the app

```bash
go run . serve
```

## 🚀 You Are All Set, Have Some Fun :D

## 🔗 REST API Endpoints

## GET ALL DATA
### GET /wsub
> Get All Data List With Preload Sublist

_Request Body_
```
not needed
```
_Request Query Params_
```
not needed
```
### GET /
> Get All Data List Without Preload Sublist

_Request Body_
```
not needed
```
_Request Query Params_
```
not needed
```

## GET DATA WITH PAGINATE
### GET /wsub/:pages
> Get Paginate Data List With Preload Sublist, in 1 page get 5 datas on it.

_Request Body_
```
not needed
```
_Request Query Params_
```
not needed
```
### GET /:pages
> Get Paginate Data List With Preload Sublist, in 1 page get 5 datas on it.

_Request Body_
```
not needed
```
_Request Query Params_
```
not needed
```

## POST DATA
### POST /list
> Post Data List, Use Form data to attach
_Request Body_
```
title : (your title)
description : (your description)
file : (attach your file)
```
_Request Query Params_
```
no need
```
### POST /sub
> Post Data Sublist, Use Body raw
_Request Body_
```
{
    "list_id" : (your list id),
    "sub_title" : (your sibb title),
    "sub_description" : (your description)
}
```
_Request Query Params_
```
no need
```

## UPDATE DATA
### PUT /list/:id
> Update Data List By id, Use Form data to attach
_Request Body_
```
title : (your title)
description : (your description)
file : (attach your file)
```
_Request Query Params_
```
no need
```

## DELETE DATA
### DELETE /:id
> Delete Data List By Id List
_Request Body_
```
no need
```
_Request Query Params_
```
no need
```
### DELETE /sub/:id
> Delete Data Sublist By Id Sublist
_Request Body_
```
no need
```
_Request Query Params_
```
no need
```

## FIND SUBLIST BY LIST ID
### GET /cart/name
> Search Data Sublist With List ID
_Request Body_
```
no need
```
_Request Query Params_
```
no need
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Gin-Gonic](https://gin-gonic.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)
