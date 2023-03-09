package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	connectToDb()

	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/task/:id", getTask)
	router.POST("/task", createTask)
	router.Run("localhost:8080")
}

func getTasks(c *gin.Context) {
	tasks, err := readTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error")
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTask(c *gin.Context) {
	id := c.Param("id")

	if idInt, err := strconv.Atoi(strings.TrimSpace(id)); err == nil {
		task, err := readTask(idInt)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error")
			return
		}
		c.IndentedJSON(http.StatusOK, task)
		return
	} else {
		fmt.Println("Invalid Id")
	}
	c.IndentedJSON(http.StatusBadRequest, "Invalid request")
}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		fmt.Println("Error decoding task json")
		return
	}

	t, err := addTask(newTask)
	if err != nil {
		fmt.Println("Error saving task")
	}
	c.IndentedJSON(http.StatusCreated, t)
}
