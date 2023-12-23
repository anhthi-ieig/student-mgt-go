# go-service-starter

### Install dependencies
```sh
go mod download
```

### Run
```sh
go run ./
```

Verify
```sh
curl http://localhost:8080/students
```

## Run database

```sh
docker-compose up
```

# Project Structure


```
.
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go                       <-- Entry point
├── pkg                           <-- packages dir
│   ├── application               <-- REST API, interfaces
│   ├── data-access               <-- Data access layer, db queries
│   ├── service                   <-- Business services
│   └── utils                     <-- Utilities
└── sql                           <-- Flyway migration scripts
    └── V1__initial_tables.sql
```

![](./docs/diagram.jpg)


## Authorization
Using JWT token for authorization and bcrypt for encrypt password
### How to setup
- Manual add 1 user into user table
- Need encode password via: https://bcrypt-generator.com/
- Use postman for checking login method (url: http:localhost:8080/login, method: POST, Authorization: Basic Auth)
- Trigger login API will return Jwt token
### How to add jwt token into request postman
- Open postman -> Add API you want to check (Ex: /students) -> select Authorization -> select Bearer -> paste token get from login API into the input.
## Middleware
- Refer example in: student-mgt-go/main.go -> GET /students API