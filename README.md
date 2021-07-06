# Elasticsearch-library

## Developers

- Chabha Mahfoufi
- Christophe Charlebois
- Guillaume Cornet

## Table of Contents

- [Directory Structure](#directory-structure)
- [Used Package](#used-package)
- [Installation](#installation)
  
## Directory Structure

```
<Wiki-Go>/
├─ api/
|   └─ server/
|       └─ routes.go
|       └─ server.go
├─ controllers/
├─ documentation/
|   └─ swagger.json
├─ middleware/
├─ models/
├─ utils/
├─ view/
├─ docker-compose.yml
├─ main.go
└─ README.md
```

## Used Package

* [mux](https://github.com/gorilla/mux) - HTTP request router and dispatcher
* [CORS](https://github.com/rs/cors) - Go cors handler

## Api server
`http://localhost:8080`

## Installation

* Launch api
``` bash
docker-compose up --build
```
