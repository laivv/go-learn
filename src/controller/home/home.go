package home

import (
	"github.com/gin-gonic/gin";
	"db"
	"model"
)

func Index(ctx *gin.Context){
	user := model.User{
		Account:"lingluo",
		Password: "123",
		Alias: "零落",
		Status: 1,
		}
	 mDb := db.Open()
	 mDb.NewRecord(user)
	 mDb.Create(&user)

	ctx.JSON(200,gin.H{
		"code":0,
		"msg":"成功",
		"data":[0]string{},
	})
}

func List(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"code":0,
		"msg":"成功",
		"data":[0]int{},
	})
}
