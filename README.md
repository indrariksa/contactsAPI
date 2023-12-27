
# Kontak Fiber (Go) and MongoDB REST API

An API built with Fiber and MongoDB.


## API Usage
Main URL / Endpoint :
https://ulbicontacts-59a677aa5bea.herokuapp.com/

#### Get all Contact

```http
  GET /contacts
```

#### Get by id Contact

```http
  GET /contacts/{id}
```

#### Add Contact

```http
  POST /insert
```


```json
{
  "nama_kontak": "Si Ujang",
  "nomor_hp": "68123456789"
}

```

#### Edit Contact

```http
  PUT /update/{id}
```


```json
{
  "nama_kontak": "Si Asep",
  "nomor_hp": "61235156431231"
}

```


#### Delete Contact

```http
  DELETE /delete/{id}
```


#### Login

```http
  POST /login
```

```json
{
  "username": "qwerty",
  "password": "qwerty"
}

```
