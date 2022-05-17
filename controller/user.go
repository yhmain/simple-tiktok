package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"admin123456": {
		Id:            2,
		Name:          "admin",
		FollowCount:   13,
		FollowerCount: 51,
		IsFollow:      true,
	},
}

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

//用户Token结构体
//JWT:Json Web Token
type UserToken struct {
	Name     string
	Password string
}

//用户注册函数，路由
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//用户鉴权token，Json序列化
	jsonBytes, _ := json.Marshal(&UserToken{Name: username, Password: password})
	token := string(jsonBytes)

	//首先判断该用户名是否已存在，为了确保用户名是唯一的
	if _, exist := service.SelectUserByName(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
			Token:    token,
		})
	} else {
		//生成新用户的ID
		newUserID := service.GetNewUserID()
		if newUserID == -1 {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: -2, StatusMsg: "生成新用户ID出错！"},
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
		err := service.InsertNewUser(newUser)
		mu.Unlock()
		//若返回空，则插入数据的时候出现错误
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: -3, StatusMsg: "创建新用户出错！"},
				Token:    token,
			})
			return
		}
		//否则，返回成功！！！
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUserID,
			Token:    token,
		})
	}
}

//用户登录函数，路由
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//用户鉴权token，Json序列化
	jsonBytes, _ := json.Marshal(&UserToken{Name: username, Password: password})
	token := string(jsonBytes)

	//检测 用户名和密码是否正确
	if user, exist := service.SelectUserByNamePwd(username, password); exist {
		fmt.Println("查找成功了呀！")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0}, //成功找到
			UserId:   user.Id,
			Token:    token,
		})
	} else { //用户名或者密码错误
		fmt.Println("用户名或者密码错误！！")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: -1, StatusMsg: "User doesn't exist"},
			Token:    token,
		})
	}
}

//输入为用户id和鉴权token，获取该用户信息
func UserInfo(c *gin.Context) {
	paramID := c.Query("user_id")  //用户id
	paramToken := c.Query("token") //鉴权token

	userID, err := strconv.ParseInt(paramID, 10, 64) //string转化为int64
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "非法的用户ID！"},
		})
		return
	}

	//用户Token 反序列化
	var token UserToken
	json.Unmarshal([]byte(paramToken), &token)

	//调用service获取数据
	//利用token中的用户名和密码获取用户ID，与传参的用户ID比较
	if user, _ := service.SelectUserByNamePwd(token.Name, token.Password); user.Id == userID {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
