# Restful GIN
Rest API with Golang, MongoDB

# Feature
* CRUD API
* Authentication
* Authorization
    * Branch `master` for role authorization
    * Branch `rbac` for dynamic role base access control authorization
* CORS
* Auto generate Swagger Docs with annotations

# Technologies
* [Gin](https://github.com/gin-gonic/gin)
* [MongoDB](https://www.mongodb.com)

# Set up
* Create file .env
* Set MongoDB URI and DB
  - PORT = "8585" or your port
  - MONGO_HOST = "your host/ localhost:27017"
  - MONGO_DB_NAME = "your db name"
  
* If you want to use real-time firebase's database. Replace with your serviceAccountKey.json. Then, add these variable into .env
  - FIREBASE_DATABASE = "your database url"
  - FIREBASE_STORAGE = "your firebase storage"

# Run
* `go mod download` for download dependencies
* `go run main.go`

# Swagger
* `localhost:8585/swagger/index.html`

