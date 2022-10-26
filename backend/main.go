package main

import (
	"backend/db"
	_ "backend/docs"
	"backend/routers"
)

// @title          Information System \"Post\"
// @version        1.0
// @description    Information System "Post" API
// @termsOfService http://swagger.io/terms/

// @contact.name  R0st0k
// @contact.email 2002rostok@gmail.com

// @host     localhost:8080
// @BasePath /api/v1

func main() {
	db.Init()

	router := routers.CreateRouters()
	router.Run(":8080")
}
