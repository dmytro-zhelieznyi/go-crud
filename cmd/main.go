package main

import (
	"crud/database"
	"crud/database/dao"
	"crud/database/model"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "go-crud"
)

var connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func main() {
	cp, err := database.GetConnectionPool(connectionString)
	if err != nil {
		log.Panic(err)
	}

	userDao := &dao.UserDao{
		ConnPool: cp,
	}

	fmt.Println("FIND ALL USERS")
	fmt.Println("----------------------------")
	users, err := userDao.FindAll()
	if err != nil {
		log.Panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println()
	fmt.Println("FIND USER BY ID")
	fmt.Println("----------------------------")
	userById, err := userDao.FindById(1)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(userById)

	fmt.Println()
	fmt.Println("CREATE USER")
	fmt.Println("----------------------------")
	createUser := &model.User{Email: "email@example.com", Age: 33}
	createDb, err := userDao.Create(createUser)
	create := createDb.(*model.User)
	fmt.Println("create:", create)

	fmt.Println()
	fmt.Println("UPDATE USER")
	fmt.Println("----------------------------")
	updateUser := &model.User{Id: create.Id - 1, Email: "update@example.com", Age: 23}
	updateDb, err := userDao.Update(updateUser)
	update := updateDb.(*model.User)
	fmt.Println("update:", update)

	fmt.Println()
	fmt.Println("DELETE USER")
	fmt.Println("----------------------------")
	_ = userDao.Delete(create.Id - 2)

	users, err = userDao.FindAll()
	if err != nil {
		log.Panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
