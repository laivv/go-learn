package model

type User struct {
  Id int `gorm:"AUTO_INCREMENT"`
  Account string
  Password string
  Alias string
  Status int
}

type UserInfo struct {
  Id int
  Account string
  CreateTime int
  LastLogin int
  ThisLogin int
}
