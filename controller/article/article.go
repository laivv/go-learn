package article

import (
	"github.com/gin-gonic/gin";
  "model"
  // "util"
  "strconv"
  // "time"
)

func Add (ctx *gin.Context){
  // util.Log("aaaa")
}

func Update(ctx *gin.Context){

}

func Delete (ctx *gin.Context){

}

func FindByPage(ctx *gin.Context){
  page , _  := strconv.Atoi(ctx.Param("page"))
  var size int = 10
  var offset int = (page - 1) * size
  var articles []model.Article
  model.InitDB()
  model.DB.Offset(offset).Limit(size).Find(&articles)
  ctx.JSON(200,gin.H{
    "code":0,
    "msg" :"成功",
    "data": articles,
  })
}
