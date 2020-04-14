package app

import (
	"github.com/ta-taiyo/bookstore-users-api/controllers/ping"
	"github.com/ta-taiyo/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
