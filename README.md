# Restful GIN
Rest API with Golang, MongoDB

# Feature
* CRUD API
* Authentication, Authorization

# Technologies
* [Gin](https://github.com/gin-gonic/gin)
* [MongoDB](https://www.mongodb.com)

# Set up
* Create file .env
* Set MongoDB URI and DB
  - MONGO_HOST = "your host/ localhost:27017"
  - MONGO_DB_NAME = "your db name"
  
* If you want to use real-time firebase's database. Replace with your serviceAccountKey.json

# Run
* `go mod download` for download dependencies
* `go run main.go`

# Swagger
* `localhost:8585/swagger/index.html`
