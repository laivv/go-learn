package app

import (
	"strconv"
	"github.com/gin-gonic/gin"
  "controller/user"
  "controller/article"
)

type App struct {
	port int
}

func(s *App) Run(port int){
	s.port = port
	r := gin.Default();
	r.POST("/login", user.Login)
  r.GET("/articles/:page",article.FindByPage)
  r.PUT("/article",article.Add)
  r.GET("/article/:id",article.FindById)
  r.POST("/article/:id",article.Update)
  r.DELETE("/article/:id",article.Delete)
  r.GET("/category/:id",article.FindByCategory)
	r.Run((":" + strconv.Itoa(s.port)))
}
