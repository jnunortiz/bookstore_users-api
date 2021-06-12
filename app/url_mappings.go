package app

import "github.com/jnunortiz/bookstore_users-api/domain/users/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
