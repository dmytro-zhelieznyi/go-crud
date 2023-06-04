package controller

import (
	"crud/database/dao"
	"crud/database/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	userDao = dao.InitUserDao()
)

func GetUsers(c *gin.Context) {
	users, err := userDao.FindAll()
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error while parsing string to int64"})
		return
	}

	user, err := userDao.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	create, err := userDao.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user = create.(*model.User)

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	create, err := userDao.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	user = create.(*model.User)

	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := userDao.Delete(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
