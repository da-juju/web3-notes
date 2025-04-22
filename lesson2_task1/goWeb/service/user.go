package service

import (
	"goweb/dao"
	"goweb/models"
)

func InsertUser(user *models.User) int64 {
	return dao.InsertUser(user)
}
