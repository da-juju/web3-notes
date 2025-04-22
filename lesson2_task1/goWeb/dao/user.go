package dao

import (
	"goweb/global"
	"goweb/models"
)

func InsertUser(user *models.User) (affected int64) {
	SqlStr := "insert into user(username,password) values (?, ?)"
	tx := global.DB.Exec(SqlStr, user.UserName, user.Password)
	affected = tx.RowsAffected
	return
}
