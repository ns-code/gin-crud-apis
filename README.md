## gin-crud-apis README
This project demos serving golang/gin based REST CRUD APIs. 

The code follows the tutorial at: https://www.allhandsontech.com/programming/golang/web-app-sqlite-go/

This project serves as a backend for https://github.com/ns-code/ng18-maintainusers-refactored

It uses an embedded SQLLite DB. The follwing APIs are currently supported:

GET /api/users

### Install
After cloning this repo, download and install https://jmeubank.github.io/tdm-gcc/download/ and set the PATH

Then run<br>
go env -w CGO_ENABLED=1

### Testing
The apis can be tested by running:

go test ./main

### Usage Notes
The gin web application can be run using:<br>
go run ./main
