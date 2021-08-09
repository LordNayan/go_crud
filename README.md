# go_crud
A normal CRUD using GO and GIN

Postman Collection - https://www.getpostman.com/collections/a8e2ca70a22fc6bbfa57

DB Used - Mongo DB (Docker)
Framework Used - Gin

Steps to run :- 

1) Boot a mongo docker image 
  
docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=myuser -e MONGO_INITDB_ROOT_PASSWORD=mypassword -e MONGO_INITDB_DATABASE=tasks -p 27017:27017 -d mongo:latest

2) Go inside the git repo and execute go_rest

./go_rest

3) Import the postman collection and try running the CRUD apis 
