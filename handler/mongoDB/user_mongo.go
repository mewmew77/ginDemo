package mongoDB

import (
	"ginDemo/common"
	"ginDemo/handler/model"
	repoModel "ginDemo/model"
	"ginDemo/repository/impl"
	"ginDemo/repository/infra"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

func NewUserMongoHandler(r gin.IRouter) {
	g := r.Group("v2")
	g.POST("add", addUser2)
	g.GET("list", listUser2)
	g.POST("update", updateUser2)
}

func addUser2(ctx *gin.Context) {
	req := model.AddUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("addUser: bindJson failed, error = %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user := &repoModel.UserForMongo{
		ID:          primitive.NewObjectID(),
		Name:        strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
	}
	repo := impl.NewUserMongoRepo(infra.GetMongoDBClient())
	if err := repo.AddUser(ctx, user); err != nil {
		log.Printf("addUser: add user to database failed, error = %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(user))
	}
}

func listUser2(ctx *gin.Context) {
	repo := impl.NewUserMongoRepo(infra.GetMongoDBClient())
	if users, err := repo.ListUser(ctx); err != nil {
		log.Printf("ListUser: get user info from database failed, error = %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(users))
	}
}

func updateUser2(ctx *gin.Context) {
	req := model.UpdateRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("updateUser: bindJson failed, error = %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	repo := impl.NewUserMongoRepo(infra.GetMongoDBClient())
	userInfo := repoModel.UserForMongo{Description: req.Description}
	err := repo.UpdateUser(ctx, bson.M{"_id": req.ID}, &userInfo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(nil))
	}
}
