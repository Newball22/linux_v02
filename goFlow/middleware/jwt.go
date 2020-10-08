package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goFlow/utils"
	"goFlow/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

/*JWT:JSON Web Token*/

var (
	JwtKey = []byte(utils.JwtKey)
	code   int
)

type MyClaims struct {
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

//生成Token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour) //设置过期时间
	claims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginflow", //发行人
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//解析Token
func ParseToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt做成中间件比较方便
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIT
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		tokens := strings.SplitN(tokenHeader, " ", 2)  //用空格分割
		if len(tokens) != 2 && tokens[0] != "Bearer" { //按照文档固定写法
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := ParseToken(tokens[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next() //调动后续的处理函数
	}
}
