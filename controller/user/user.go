package user

import (
	"github.com/gin-gonic/gin";
  "model"
  // "time"
)



// func Index(ctx *gin.Context){
// 	user := model.User{
// 		Account:"lingluo",
// 		Password: "123",
//     Alias: "零落",
//     Type: 0,
//     CreateDate: time.Now(),
//     LastLoginDate: time.Now(),
//     CurrLoginDate: time.Now(),
//     LoginCount: 1,
//     Email:"test@test.com",
// 		State: 1,
// 	}

//   model.DB.Create(&user)
// 	ctx.JSON(200,gin.H{
// 		"code":0,
// 		"msg":"成功",
// 		"data":[0]string{},
// 	})
// }


// func List(ctx *gin.Context){
// 	ctx.JSON(200,gin.H{
// 		"code":0,
// 		"msg":"成功",
// 		"data":[0]int{},
// 	})
// }
func Login(ctx *gin.Context) {
  var user model.User
  if err := ctx.ShouldBindJSON(&user) ; err != nil {

  } else {
    FindUser(&user)
  }
}

func FindUser(user *model.User) {
  model.DB.Find(user)
}
