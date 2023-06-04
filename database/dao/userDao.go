package dao

import (
	"crud/database"
	"crud/database/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UserDao struct {
	ConnPool *database.ConnectionPool
}

func (dao *UserDao) FindAll() ([]interface{}, error) {
	db := dao.ConnPool.GetConnection()
	defer dao.ConnPool.ReleaseConnection(db)

	query := "SELECT * FROM usr"
	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()

	var users []interface{}
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.Id, &user.Email, &user.Age); err != nil {
			log.Panic(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (dao *UserDao) FindById(id int64) (interface{}, error) {
	db := dao.ConnPool.GetConnection()
	defer dao.ConnPool.ReleaseConnection(db)

	query := "SELECT * FROM usr WHERE id=$1"
	row := db.QueryRow(query, &id)

	user := &model.User{}
	if err := row.Scan(&user.Id, &user.Email, &user.Age); err != nil {
		log.Panic(err)
		return nil, err
	}

	return user, nil
}

func (dao *UserDao) Create(userI interface{}) (interface{}, error) {
	db := dao.ConnPool.GetConnection()
	defer dao.ConnPool.ReleaseConnection(db)

	user, ok := userI.(*model.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type")
	}

	query := "INSERT INTO usr (email, age) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, user.Email, user.Age).Scan(&user.Id)

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return user, nil
}

func (dao *UserDao) Update(userI interface{}) (interface{}, error) {
	db := dao.ConnPool.GetConnection()
	defer dao.ConnPool.ReleaseConnection(db)

	user, ok := userI.(*model.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type")
	}

	query := "UPDATE usr SET email=$1, age=$2 WHERE id=$3"
	_, err := db.Exec(query, user.Email, user.Age, user.Id)

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return user, nil
}

func (dao *UserDao) Delete(id int64) error {
	db := dao.ConnPool.GetConnection()
	defer dao.ConnPool.ReleaseConnection(db)

	query := "DELETE FROM usr WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
