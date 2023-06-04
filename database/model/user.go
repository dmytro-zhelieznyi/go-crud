package model

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}
