package model
import (
  "github.com/jinzhu/gorm"
  "util"
  _ "github.com/go-sql-driver/mysql"
  "time"
  "fmt"
  "config"
)

var DB *gorm.DB
var dbConfig = config.DBConfig
var url string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["database"], dbConfig["charset"])

func InitDB() {
  if DB != nil {
    return
  }
  db,err := gorm.Open(dbConfig["type"].(string), url)
  if err == nil {
    util.Log("数据库连接失败")
  } else {
    util.Log("数据库连接成功")
  }
  // defer db.Close()
  // 使用表名单数形式
  db.SingularTable(true)
  DB = db
}



type User struct {
  ID int
	Account string
	Password string
	Alias string
	Type uint
	CreateDate time.Time
	LastLoginDate  time.Time
	CurrLoginDate time.Time
	LoginCount int
	Email string
	State uint
}

