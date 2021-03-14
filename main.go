package main

import users "github.com/yarikyarichek/user-service"

func main() {
	srv := new(users.UserService)
	srv.Start()
}
