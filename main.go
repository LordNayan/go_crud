package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//Function to get env variables from .env
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func handleCreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleGetTasks(c *gin.Context) {
	var task Task
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error in reading json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		log.Printf("Error in unmarshalling json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	result, error := GetByTitle(task.Title)
	if error != nil {
		log.Printf("Error in mongo - GET %v", error)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func handleUpdateTask(c *gin.Context) {
	var task Task
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error in reading json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		log.Printf("Error in unmarshalling json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	error := UpdateByID(&task)
	if error != nil {
		log.Printf("Error in mongo - DELETE %v", error)
		c.JSON(http.StatusNotFound, gin.H{"msg": error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": 1})
}

func handleDeleteTask(c *gin.Context) {
	var task Task
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error in reading json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		log.Printf("Error in unmarshalling json Body %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	error := DeleteByID(task.ID)
	if error != nil {
		log.Printf("Error in mongo - DELETE %v", error)
		c.JSON(http.StatusNotFound, gin.H{"msg": error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": 1})
}

func main() {
	//Default Route
	r := gin.Default()

	//CRUD Routes
	r.POST("/tasks/", handleCreateTask)
	r.GET("/tasks/", handleGetTasks)
	r.PUT("/tasks/", handleUpdateTask)
	r.DELETE("/tasks/", handleDeleteTask)

	//START
	r.Run()
}
