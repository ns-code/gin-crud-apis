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
The apis can be unit tested by running:<br>
go test .

The apis can be integration tested by running the swagger UI:<br>
http://localhost:8080/docs/index.html

### Usage Notes
The gin web application can be run using:<br>
go run .

and then accessing:<br>
http://localhost:8080/api/users
