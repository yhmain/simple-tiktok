package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yhmain/simple-tiktok/model"
)

//NOTE: 以下为基于jwt-go实现的token权限认证

// 用户Token结构体  JWT:Json Web Token
type UserToken struct {
	UserID   int64
	Name     string
	Password string
}

//用于生成Token的结构体
type UserClaims struct {
	UserToken
	jwt.RegisteredClaims //v4版本新增
}

var jwtKey = []byte("tiktok")             //定义Secret
const TokenExpireDuration = time.Hour * 2 //定义JWT的过期时间：2小时

//发放Token
func GenToken(userToken *UserToken) (string, error) {
	// 创建一个用户的声明，即初始化结构体 UserClaims
	c := UserClaims{
		UserToken: *userToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c) // 使用指定的签名方法创建签名对象
	return token.SignedString(jwtKey)                     // 使用指定的Secret签名并获得完整的编码后的字符串token
}

//解析Token
func ParseToken(tokenString string) (*jwt.Token, *UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// 自定义函数：JWTAuthUser 基于JWT的认证中间件
func JWTAuthUserToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取token，并解析
		token := c.Query("token")
		_, claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, InvalidTokenErr) //token解析失败
			c.Abort()
			return
		}
		//获取user_id，并与token解析出来的进行对比
		//在获取publish list时，本来就没有user_id !!!!!
		paramID := c.Query("user_id")
		if paramID != "" && strconv.FormatInt(claims.UserID, 10) != paramID {
			// fmt.Printf("user_id: %v\n", paramID)
			// fmt.Printf("token_user_id: %v\n", strconv.FormatInt(claims.UserID, 10))
			// fmt.Printf("%v\n", ValidateTokenErr)
			c.JSON(http.StatusOK, ValidateTokenErr) //token校验失败
			c.Abort()
			return
		}

		c.Set("usertoken", claims.UserToken)
		c.Next() // 执行后续的处理函数
	}
}

//还在测试中
// 自定义函数：JWTAuthPublishAction 基于JWT的认证中间件
func JWTAuthPublishAction() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.MustGet("token").(string) //从上下文中获取已经保存的token
		publishToken := c.Query("token")     //获取token
		if publishToken != token {
			c.JSON(http.StatusOK, ValidateTokenErr.Error())
			c.Abort()
			return
		}
		user := c.MustGet("user").(model.User) //从上下文中获取已经保存的user
		title := c.Query("title")              //获取 title
		fmt.Printf("%v \n%v\n", user, title)
		c.Abort()
		// c.Next() // 去执行后续的处理函数
	}
}
