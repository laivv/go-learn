package article

import (
	"github.com/gin-gonic/gin";
  "model"
  "math"
  // "math/big"
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
  page , err  := strconv.Atoi(ctx.Param("page"))
  if err != nil || page < 1 {
    page = 1
  }
  var size int = 10
  var offset int = (page - 1) * size
  var articles []model.Article
  var dataCount int = 0
  model.InitDB()
  model.DB.Model(&model.Article{}).Count(&dataCount)
  var pageCount = int(math.Ceil(float64(dataCount) / float64(size)))
  if page > pageCount {
    page = pageCount
  }
  model.DB.Offset(offset).Limit(size).Find(&articles)
  ctx.JSON(200,gin.H{
    "code":0,
    "msg" :"成功",
    "data": map[string]interface{} {
      "list":articles,
      "pageIndex":page,
      "pageSize":size,
      "pageCount":pageCount,
      "dataCount":dataCount,
    },
  })
}
