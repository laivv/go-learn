package util

import (
	"strconv"
	"reflect"
  "fmt"
  "github.com/gin-gonic/gin"
)

func typeOf(arg interface{}) string{
	return reflect.TypeOf(arg).String()
}

func toString(arg interface{}) string{
	ret := ""
	if typeOf(arg) == "int" {
		// 类型断言
		ret = strconv.Itoa(arg.(int))
	}
	if typeOf(arg) == "string" {
		// 类型断言
		ret = arg.(string)
	}
	return ret
}

func Log(args ...interface{}){
	str := ""
	for _,arg := range args{
		str += toString(arg)
	}
	fmt.Println(str)
}

const (
  MESSAGE_ERROR = "失败"
  MESSAGE_SUCCESS = "成功"
  CODE_SUCCESS = 0
  CODE_ERROR = 1
)

func Send(ctx *gin.Context, args ...interface{}) {
  httpCode := 200
  code := CODE_SUCCESS
  msg := MESSAGE_SUCCESS
  var data interface{} = nil
  for index,arg := range args{
    if index == 0 {
      if typeOf(arg) == "int" {
        httpCode = arg.(int)
        if httpCode != 200 && httpCode != 304 {
          msg = MESSAGE_ERROR
          code = CODE_ERROR
        }
      } else {
        data = arg
      }
    } else if index == 1 {
      data = arg
    } else if index == 2 && typeOf(arg) == "string" {
      msg = arg.(string)
    }
  }
  ctx.JSON(httpCode, gin.H{
    "code": code,
    "msg": msg,
    "data": data,
  })
}
