package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

//NOTE: 以下为基于jwt-go实现的token权限认证

// 用户Token结构体  JWT:Json Web Token
type UserToken struct {
	Name     string
	Password string
}

type UserClaims struct {
	UserName             string
	Password             string
	jwt.RegisteredClaims //v4版本新增
}

var jwtKey = []byte("tiktok")             //定义Secret
const TokenExpireDuration = time.Hour * 2 //定义JWT的过期时间：2小时

//发放Token
func GenToken(UserToken *UserToken) (string, error) {
	// 创建一个用户的声明，即初始化结构体 UserClaims
	c := UserClaims{
		UserName: UserToken.Name,     //用户输入的账号
		Password: UserToken.Password, //用户输入的密码
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
func JWTAuthUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		paramID := c.Query("user_id")                    //获取user_id
		userID, err := strconv.ParseInt(paramID, 10, 64) //string转化为int64
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: InvalidUserIDErr, //非法的用户ID
			})
			c.Abort()
			return
		}
		token := c.Query("token") //获取token
		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析它
		_, mc, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, InvalidTokenErr.Error())
			c.Abort()
			return
		}
		//调用service获取数据
		//利用token中的用户名和密码获取用户ID，与传参的用户ID比较
		var user model.User
		if user, _ := service.SelectUserByNamePwd(mc.UserName, mc.Password); user.Id != userID {
			c.JSON(http.StatusOK, UserResponse{
				Response: UserNotExistErr, //用户不存在
			})
			c.Abort()
			return
		}
		// 将当前请求的user保存到请求的上下文c上
		c.Set("user", user)
		c.Next() // 后续的处理函数可以用c.Get("username")来获取当前请求的用户信息
	}
}

// 自定义函数：JWTAuthUser 基于JWT的认证中间件
func JWTAuthPublishList() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token") //获取token
		_, mc, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, InvalidTokenErr.Error())
			c.Abort()
			return
		}
		//调用service获取数据
		//利用token中的用户名和密码获取用户ID
		user, exist := service.SelectUserByNamePwd(mc.UserName, mc.Password)
		if !exist {
			c.JSON(http.StatusOK, UserResponse{
				Response: UserNotExistErr, //用户不存在
			})
			c.Abort()
			return
		}
		// 将当前请求的user保存到请求的上下文c上
		c.Set("user", user)
		c.Next() // 后续的处理函数可以用c.Get("username")来获取当前请求的用户信息
	}
}
