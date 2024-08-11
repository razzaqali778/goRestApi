package controllers

import (
	"users.com/daos"
	"users.com/utils"
)

type User struct {
	utils   utils.Utils
	userDAO daos.User
}
