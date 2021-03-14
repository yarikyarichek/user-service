package user

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name        string
	Email       string
	DateOfBirth time.Time
}

type UserDB map[string]*User

type UserService struct {
	db *UserDB
}

var localDB = UserDB{
	"1": {
		Name:        "Stepan",
		Email:       "test1@gmail.com",
		DateOfBirth: time.Date(1990, 1, 21, 0, 0, 0, 0, time.UTC),
	},
	"2": {
		Name:        "Boris",
		Email:       "test2@gmail.com",
		DateOfBirth: time.Date(2000, 7, 9, 0, 0, 0, 0, time.UTC),
	},
	"3": {
		Name:        "Nikola",
		Email:       "test3@gmail.com",
		DateOfBirth: time.Date(2010, 7, 1, 0, 0, 0, 0, time.UTC),
	},
}

// host:port/users
func (srv *UserService) getUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, srv.db.String())
}

// host:port/greetings
func (srv *UserService) getGreetings(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, world!\n")
}

func (srv *UserService) Start() error {
	srv.db = &localDB
	http.HandleFunc("/users", srv.getUsers)
	http.HandleFunc("/greetings", srv.getGreetings)
	return http.ListenAndServe("127.0.0.1:8080", nil)
}

func (user *User) String() string {
	return fmt.Sprintf("Name: %s, Email: %s, Age: %d", user.Name, user.Email, int(time.Now().Sub(user.DateOfBirth).Hours()/24/365))
}

func (db *UserDB) String() string {
	var resalt string
	for _, user := range *db {
		resalt = fmt.Sprintf("%s%s\n", resalt, user.String())
	}
	return resalt
}
