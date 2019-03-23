  package db

  import (
    "github.com/jinzhu/gorm"
    "util"
    _ "github.com/go-sql-driver/mysql"
  )

  func Open() *gorm.DB{
    db,err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
    if err == nil {
      util.Log("数据库连接失败")
    }else{
      util.Log("数据库连接成功")
    }
    // defer db.Close()
    return db
  }

