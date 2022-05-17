// 该包定义各种error对应的常量及内容说明
package controller

import (
	"errors"
	"fmt"
)

//常量  Status Code
const (
	SuccessCode         = 0
	ServiceErrCode      = 10001
	UserNotExistErrCode = 10004

	UserAlreadyExistErrCode = 10005
	GenNewUserIDErrCode     = 10006
	InsertNewUserCode       = 10007
	InvalidUserIDCode       = 10008

	AuthHeaderEmptyCode  = 20001
	AuthHeaderFormatCode = 20002
	InvalidTokenCode     = 20003
	GenTokenFailedCode   = 20004
	ValidateTokenCode    = 20005
)

//供其他.go文件使用自定义报错信息
var (
	Success         = NewErrNo(SuccessCode, "Success")
	ServiceErr      = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	UserNotExistErr = NewErrNo(UserNotExistErrCode, "用户不存在")

	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "用户名已经存在")
	GenNewUserIDErr     = NewErrNo(GenNewUserIDErrCode, "生成新用户ID出错")
	InsertNewUserErr    = NewErrNo(InsertNewUserCode, "插入新用户出错")
	InvalidUserIDErr    = NewErrNo(InvalidUserIDCode, "非法的用户ID")

	AuthHeaderEmptyErr  = NewErrNo(AuthHeaderEmptyCode, "请求头中Auth为空")
	AuthHeaderFormatErr = NewErrNo(AuthHeaderFormatCode, "请求头中Auth格式有误")
	InvalidTokenErr     = NewErrNo(InvalidTokenCode, "无效的Token字符串")
	GenTokenFailedErr   = NewErrNo(GenTokenFailedCode, "鉴权Token生成失败")
	ValidateTokenErr    = NewErrNo(GenTokenFailedCode, "Token校验失败")
)

//响应结构体
type Response struct {
	StatusCode int    `json:"status_code"`          //状态码
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}

//返回一个错误信息的字符串
func (e Response) Error() string {
	return fmt.Sprintf("status_code=%d, status_msg=%s", e.StatusCode, e.StatusMsg)
}

//支持自定义一个ErrNo结构体
func NewErrNo(code int, msg string) Response {
	return Response{code, msg}
}

////支持自定义一个ErrNo结构体（不带code）
func (e Response) WithMessage(msg string) Response {
	e.StatusMsg = msg
	return e
}

// ConvertErr convert error to Errno（把系统的error类型转化成自定义的ErrNo结构体类型）
func ConvertErr(err error) Response {
	Err := Response{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.StatusMsg = err.Error()
	return s
}
