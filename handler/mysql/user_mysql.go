package mysql

import (
	"ginDemo/common"
	"ginDemo/handler/model"
	repoModel "ginDemo/model"
	"ginDemo/repository/impl"
	"ginDemo/repository/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewUserMysqlHandler(r gin.IRouter) {
	g := r.Group("v1")
	g.POST("add", addUser)
	g.GET("list", listUser)
}

func addUser(ctx *gin.Context) {
	req := model.AddUserRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user := &repoModel.UserForMysql{
		Name:        strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
	}
	repo := impl.NewUserMysqlRepo(infra.GetMysqlDB())
	if err := repo.AddUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(user))
	}
}

func listUser(ctx *gin.Context) {
	repo := impl.NewUserMysqlRepo(infra.GetMysqlDB())
	if users, err := repo.ListUser(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(users))
	}
}
