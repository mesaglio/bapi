package controllers

import (
	"bapi/go-gin-server/src/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var users []*models.User
var mutex sync.Mutex

func init() {
	users = []*models.User{}
	mutex = sync.Mutex{}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	user := getUserFromBody(c)
	if user != nil {
		mutex.Lock()
		users = append(users, user)
		mutex.Unlock()
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	mutex.Lock()
	user := getUserByUsername(username)
	mutex.Unlock()
	if user != nil && username != "" {
		c.JSON(http.StatusOK, user)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func DeleteUserByUsername(c *gin.Context) {
	username := c.Param("username")
	mutex.Lock()
	deleteUserByUsername(username)
	mutex.Unlock()
	c.Status(http.StatusOK)
}

func UpdateUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user := getUserFromBody(c)
	if user != nil && username == user.Username {
		mutex.Lock()
		deleteUserByUsername(username)
		users = append(users, user)
		mutex.Unlock()
		c.JSON(http.StatusOK, user)
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func deleteUserByUsername(username string) {
	if username == "" {
		return
	}
	for i, elem := range users {
		if isTheUser(username, elem) {
			users[i] = users[len(users)-1]
			users[len(users)-1] = nil
			users = users[:len(users)-1]
			return
		}
	}
}

func isTheUser(username string, elem *models.User) bool {
	return elem.Username == username
}

func getUserByUsername(username string) *models.User {
	for _, elem := range users {
		if isTheUser(username, elem) {
			return elem
		}
	}
	return nil
}

func getUserFromBody(c *gin.Context) *models.User {
	user := models.User{}
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("getUserFromBody: %v", err)
		return nil
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Printf("getUserFromBody: %v", err)
		return nil
	}
	if user.Username == "" || user.Email == "" {
		return nil
	}
	return &user
}
