package router

import (
	"github.com/gin-gonic/gin";
)




func Index(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"code":0,
		"msg":"成功",
		"data":[...]string{},
	})
}

func List(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"code":0,
		"msg":"成功",
		"data":[...]int{},
	})
}