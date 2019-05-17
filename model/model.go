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
var url string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["database"], dbConfig["charset"])

func init() {
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
   ID uint `json:"id"`
	 Title string `json:"title"`
	 Content string `json:"content"`
	 CategoryId uint `json:"categoryId"`
	 CreateDate string `gorm:"TYPE:DATETIME" json:"createDate"`
	 UpdateDate string `json:"updateDate"`
	 DeleteDate string `json:"deleteDate"`
	 CreateUser int `json:"createUser"`
	 VisitCount uint `json:"visitCount"`
	 CommentState uint `json:"commentState"`
	 State uint `json:"state"`
	 Permission uint `json:"permission"`
	 Reminder string `json:"reminder"`
	 ReminderKey string `json:"reminderKey"`
}

type User struct {
  ID uint `json:"id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Alias string `json:"alias"`
	Type uint `json:"type"`
	CreateDate time.Time `json:"createDate"`
	LastLoginDate  time.Time `json:"lastLoginDate"`
	CurrLoginDate time.Time `json:"currLoginDate"`
	LoginCount int `json:"loginCount"`
	Email string `json:"email"`
	State uint `json:"state"`
}

type Search struct {
  ID uint `json:"id"`
	Name string `json:"name"`
	Count uint `json:"count"`
	Type uint `json:"type"`
	CreateDate  time.Time `json:"createDate"`
	LastSearchDate time.Time `json:"lastSearchDate"`
}

type Message struct {
  ID uint `json:"id"`
	UserName string `json:"userName"`
	Content string `json:"content"`
	Ipv4 string `json:"ipv4"`
	Email string `json:"email"`
	CreateDate time.Time `json:"createDate"`
	UserType  uint `json:"userType"`
	State uint `json:"state"`
}

