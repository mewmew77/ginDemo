package main

import (
	"ginDemo/handler/basic"
	"ginDemo/handler/mongoDB"
	"ginDemo/handler/websocket"
	"github.com/gin-gonic/gin"
)

func initModules(r gin.IRouter) {
	basic.NewBasicHandler(r)
	//handler.NewUserMysqlHandler(r)
	websocket.NewWSHandler(r)
	mongoDB.NewUserMongoHandler(r)
}
