package model
import (
  "time"
)

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

type UserInfo struct {
  ID int
  Account string
  CreateTime int
  LastLogin int
  ThisLogin int
}
