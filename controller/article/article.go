package article

import (
	"github.com/gin-gonic/gin"
  "model"
  "math"
  // "math/big"
  "util"
  "strconv"
  // "time"
)

func Add(ctx *gin.Context){
  var article model.Article
  if err := ctx.ShouldBindJSON(&article) ; err != nil {
    util.Send(ctx,400)
    return
  }
  if er := model.DB.Create(&article).Error ; er != nil {
    util.Send(ctx,500)
  } else {
    util.Send(ctx)
  }
}

func Update(ctx *gin.Context){
  var article model.Article
  if err := ctx.ShouldBindJSON(&article) ; err != nil {
    util.Send(ctx,400)
    return
  }
 if er := model.DB.Update(&article) ; er != nil {
   util.Send(ctx,500)
 } else {
   util.Send(ctx)
 }
}

func Delete(ctx *gin.Context){
  id , err := strconv.Atoi(ctx.Param("id"))
  sendError := func (){
    util.Send(ctx,400,nil,"无效的参数：id")
  }
  if err != nil || id < 0 {
    sendError()
    return
  }
  article := model.Article{
    ID:uint(id),
  }
  if model.DB.Delete(&article).Error != nil {
    sendError()
  } else {
    util.Send(ctx)
  }
}

func FindById(ctx *gin.Context){
  id , err  := strconv.Atoi(ctx.Param("id"))
  if err != nil || id < 0 {
    util.Send(ctx,400,nil,"无效的参数：id")
    return
  }
  article := model.Article{
    ID:uint(id),
  }
  if model.DB.Find(&article).Error != nil {
    util.Send(ctx,404)
  } else {
    util.Send(ctx,article)
  }
}

func FindByPage(ctx *gin.Context){
  page, err  := strconv.Atoi(ctx.Param("page"))
  var articles []model.Article
  var dataCount int = 0
  var size int = 10
  model.DB.Model(&model.Article{}).Count(&dataCount)
  var pageCount = int(math.Ceil(float64(dataCount) / float64(size)))
  if err != nil || page < 1 {
    page = 1
  }
  if page > pageCount {
    page = pageCount
  }
  var offset int = (page - 1) * size
  model.DB.Offset(offset).Limit(size).Find(&articles)
  util.Send(ctx, map[string]interface{} {
    "list":articles,
    "pageIndex":page,
    "pageSize":size,
    "pageCount":pageCount,
    "dataCount":dataCount,
  })
}

func FindByCategory(ctx *gin.Context){
  _,err := strconv.Atoi(ctx.Param("id"))
  if err != nil {
    util.Send(ctx,400)
  }
}
