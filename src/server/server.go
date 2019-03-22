package server

import (
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"router"
	// "log"
)


type Runable interface {
	Run()
}
type Server struct {
	port int
}



func(s *Server) Run(port int){
	s.port = port


	db,err := gorm.Open("localhost","blog")
	if err == nil {
		// log.Log("数据库连接失败")
	}
	if db != nil{
		// log.Log("数据库连接成功")
	}




	r := gin.Default();
	r.GET("/", router.Index)
	r.GET("/list", router.List)
	p := ":" + strconv.Itoa(s.port)
	r.Run(p)

	// log.Log("server is running on port",port)

}