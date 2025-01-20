## gin-crud-apis README
This project demos serving golang/gin based REST CRUD APIs. 

The code follows the tutorial at: https://www.allhandsontech.com/programming/golang/web-app-sqlite-go/

This project serves as a backend for https://github.com/ns-code/ng18-maintainusers-refactored

It uses an embedded SQLLite DB. The following APIs are currently supported:

GET /api/users
POST /api/users
PUT /api/users/{userId}
DELETE /api/users/{userId}


### Releases Info
The Release-1.0.0 uses gin and std database/sql packages. The Release-2.0.0 uses gorilla/mux and sqlx packages


### Install
After cloning this repo, download and install https://jmeubank.github.io/tdm-gcc/download/ and set the PATH

Then run<br>
go env -w CGO_ENABLED=1


### Unit Testing
The apis can be unit tested by running:<br>
go test .


### Integration Testing
Start the web application using:<br>
go run .

All the APIs mentioned above can then be integration tested using the Swagger UI:<br>
http://localhost:8080/docs/index.html
