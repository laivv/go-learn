package app

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"controller/home"
	"util"
)


type Runable interface {
	Run()
}
type App struct {
	port int
}

func(s *App) Run(port int){
	s.port = port
	r := gin.Default();
	r.GET("/", home.Index)
	r.GET("/list", home.List)
	r.Run((":" + strconv.Itoa(s.port)))
	util.Log("server is running on port ",port)

}
