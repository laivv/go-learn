package main

import (
	"fmt"
	"github.com/jinzhu/gorm";
	"github.com/gin-gonic/gin";
)

func main (){
	server();
}

func server(){
	db,err := gorm.Open("localhost","blog")
	if err == nil {
		fmt.Println("数据库连接失败")
	}
	if db != nil{
		fmt.Println("数据库连接成功")
	}
	port  := 8001
	router := gin.Default();
	router.GET("/",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{"msg":"成功"})
	})
	router.Run(":8001");
	fmt.Println("server is running onport %d",port)
}