// 该包定义各种error对应的常量及内容说明
package controller

import (
	"errors"
	"fmt"
)

//响应结构体
type Response struct {
	StatusCode int    `json:"status_code"`          //状态码
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}

//常量  Status Code
const (
	SuccessCode         = 0
	ServiceErrCode      = 10001
	UserNotLoginErrCode = 10002
	UserNotExistErrCode = 10003

	UserAlreadyExistErrCode = 10004
	GenNewUserIDErrCode     = 10005
	InsertNewUserCode       = 10006
	InvalidUserIDCode       = 10007
	InvalidVideoIDCode      = 10008

	AuthHeaderEmptyCode  = 20001
	AuthHeaderFormatCode = 20002
	InvalidTokenCode     = 20003
	GenTokenFailedCode   = 20004
	ValidateTokenCode    = 20005

	InvalidTimeCode        = 30001
	TooLongInputCode       = 30002
	UploadFileFailedCode   = 30003
	ExtractVideoFailedCode = 30004

	UpdateSQLCode = 50001
)

//供其他.go文件使用自定义报错信息
var (
	Success         = NewResponse(SuccessCode, "Success")
	ServiceErr      = NewResponse(ServiceErrCode, "Service is unable to start successfully")
	UserNotLoginErr = NewResponse(UserNotLoginErrCode, "用户还未登录")
	UserNotExistErr = NewResponse(UserNotExistErrCode, "用户不存在")

	UserAlreadyExistErr = NewResponse(UserAlreadyExistErrCode, "用户名已经存在")
	GenNewUserIDErr     = NewResponse(GenNewUserIDErrCode, "生成新用户ID出错")
	InsertNewUserErr    = NewResponse(InsertNewUserCode, "插入新用户出错")
	InvalidUserIDErr    = NewResponse(InvalidUserIDCode, "非法的用户ID")
	InvalidVideoIDErr   = NewResponse(InvalidVideoIDCode, "非法的视频ID或action_type")

	AuthHeaderEmptyErr  = NewResponse(AuthHeaderEmptyCode, "请求头中Auth为空")
	AuthHeaderFormatErr = NewResponse(AuthHeaderFormatCode, "请求头中Auth格式有误")
	InvalidTokenErr     = NewResponse(InvalidTokenCode, "无效的Token字符串")
	GenTokenFailedErr   = NewResponse(GenTokenFailedCode, "鉴权Token生成失败")
	ValidateTokenErr    = NewResponse(ValidateTokenCode, "Token校验失败")

	InvalidTimeErr        = NewResponse(InvalidTimeCode, "非法的时间戳")
	TooLongInputErr       = NewResponse(TooLongInputCode, "输入内容超过长度限制")
	UploadFileFailedErr   = NewResponse(UploadFileFailedCode, "文件上传失败")
	ExtractVideoFailedErr = NewResponse(ExtractVideoFailedCode, "视频封面生成失败")

	UpdateSQLErr = NewResponse(UpdateSQLCode, "数据库更新出错")
)

//返回一个错误信息的字符串
func (e Response) Error() string {
	return fmt.Sprintf("status_code=%d, status_msg=%s", e.StatusCode, e.StatusMsg)
}

//支持自定义一个Response结构体
func NewResponse(code int, msg string) Response {
	return Response{code, msg}
}

////支持自定义一个Response结构体（不带code）
func (e Response) WithMessage(msg string) Response {
	e.StatusMsg = msg
	return e
}

// ConvertErr convert error to Response（把系统的error类型转化成自定义的Response结构体类型）
func ConvertErr(err error) Response {
	Err := Response{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.StatusMsg = err.Error()
	return s
}
