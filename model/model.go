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

type Article struct {
   ID uint
	 Title string
	 Content string
	 CategoryId uint
	 CreateDate time.Time
	 UpdateDate time.Time
	 DeleteDate time.Time
	 CreateUser int
	 VisitCount uint
	 CommentState uint
	 State uint
	 Permission uint
	 Reminder string
	 ReminderKey string
}

type User struct {
  ID uint
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

type Search struct {
  ID uint
	Name string
	Count uint
	Type uint
	CreateDate  time.Time
	LastSearchDate time.Time
}

type Message struct {
  ID uint
	UserName string
	Content string
	Ipv4 string
	Email string
	CreateDate time.Time
	UserType  uint
	State uint
}

