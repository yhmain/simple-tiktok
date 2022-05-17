package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

//用户登录响应体
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//用户响应体
type UserResponse struct {
	Response
	User model.User `json:"user"`
}

//用户注册函数，路由
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//生成用户鉴权token
	token, err := GenToken(&UserToken{Name: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: GenTokenFailedErr,
			Token:    "",
		})
	}

	//首先判断该用户名是否已存在，为了确保用户名是唯一的
	if _, exist := service.SelectUserByName(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: UserAlreadyExistErr,
			Token:    token,
		})
	} else {
		//生成新用户的ID
		newUserID := service.GetNewUserID()
		if newUserID == -1 {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: GenNewUserIDErr,
				Token:    token,
			})
			return
		}
		//构造新的用户结构体
		newUser := model.User{
			Id:       newUserID,
			NickName: username,
			Password: password,
		}
		//执行插入操作，加锁
		var mu sync.Mutex
		mu.Lock()
		err := service.InsertNewUser(&newUser) //当调用一个函数时，会对其每一个参数值进行拷贝。这种情况需要用到指针
		mu.Unlock()
		//若返回空，则插入数据的时候出现错误
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: InsertNewUserErr,
				Token:    token,
			})
			return
		}
		//否则，返回成功！！！
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Success,
			UserId:   newUserID,
			Token:    token,
		})
	}
}

//用户登录函数，路由
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//生成用户鉴权token
	token, err := GenToken(&UserToken{Name: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: GenTokenFailedErr,
			Token:    "",
		})
	}

	//检测 用户名和密码是否正确
	if user, exist := service.SelectUserByNamePwd(username, password); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Success, //成功找到
			UserId:   user.Id,
			Token:    token,
		})
	} else { //用户名或者密码错误
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: UserAlreadyExistErr,
			Token:    token,
		})
	}
}

//输入为用户id和鉴权token，获取该用户信息
//注意：有中间件已处理
func UserInfo(c *gin.Context) {
	user := c.MustGet("user").(model.User)

	//此时代表成功！
	c.JSON(http.StatusOK, UserResponse{
		Response: Success,
		User:     user,
	})
}
