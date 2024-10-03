package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	createdon time.Time `json:"createdon"`
}

func GetUsers() []User {
	return []User{
		User{Id: uuid.NewString(), FirstName: "Brian", LastName: "Lee", UserName: "brian", Email: "brian@gmail.com", Password: "3423kksdf", createdon: time.Now()},
		User{Id: uuid.NewString(), FirstName: "Philip", LastName: "Pan", UserName: "philip", Email: "philip@gmail.com", Password: "215542", createdon: time.Now()},
		User{Id: uuid.NewString(), FirstName: "Aaron", LastName: "Green", UserName: "aaron", Email: "aaron@gmail.com", Password: "877554", createdon: time.Now()},
	}
}
