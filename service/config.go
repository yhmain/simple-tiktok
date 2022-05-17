package service

import (
	"bytes"
)

const (
	PREFIX_VIDEOS string = "http://192.168.1.108:8081/videos/"
	PREFIX_COVERS string = "http://192.168.1.108:8081/covers/"
)

//不定参数函数
//连接多个个字符串  s1+s2+...
func ConcatByBuilder(s ...string) string {
	var bt bytes.Buffer
	// 向bt中写入字符串
	for _, e := range s {
		bt.WriteString(e)
	}
	return bt.String()
}
