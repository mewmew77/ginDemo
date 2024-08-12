package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var (
	// 秘钥，用于对jwt进行签名
	jwtKey = []byte("gin_demo_test_key")

	// 默认token过期时间
	defaultExpireTime = 10 * time.Minute

	// 保存的token信息
	tokenInfo = CacheToken{
		tokenMap: make(map[string]string, 0),
	}
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

type CacheToken struct {
	sync.RWMutex
	tokenMap map[string]string // k-v = username-token
}

func CheckToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("token")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, "need access token")
		return
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
}

// InitToken 登录身份校验成功后使用
func InitToken(userName string) string {
	claims := &Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(defaultExpireTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	// 记录token
	tokenInfo.Lock()
	tokenInfo.tokenMap[userName] = tokenStr
	tokenInfo.Unlock()
	return tokenStr
}

// UpdateToken 保持用户操作期间不掉线
func UpdateToken(userName string) {

}
