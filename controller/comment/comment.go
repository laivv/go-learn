package comment

import (
	"github.com/gin-gonic/gin";
  "model"
  "util"
  "strconv"
)


// func FindByPage(ctx *gin.Context) {
//   if articleId ,err := strconv.Atoi(ctx.Param("id")) ; err != nil {
//     util.Send(ctx,400)
//     return
//   }
//   var comment model.Comment
//   model.DB.Offset(offset).Limit(size).Find(&comment)
// }
