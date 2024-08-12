package main

import (
	"ginDemo/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.CheckToken)
	routers := initModules
	routers(engine)
	err := engine.Run("localhost:8080")
	if err != nil {
		log.Fatalf("run server error, err = %v", err)
	}
}
