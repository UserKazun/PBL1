package service

import (
	"errors"
	"fmt"

	"github.com/PBL1/model"
	"github.com/gin-gonic/gin"
)

// CreateUser ...DBに与えられたデータをinsertする
func CreateUser(user model.User) (model.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// PostLoginDataInCookie ...ログイン時にクッキーにログイン情報を保持させる
func PostLoginDataInCookie(c *gin.Context, userID string, password string) (*model.User, error) {
	User := model.User{}

	err := db.Where("id = ? and password = ?", userID, password).First(&User).Error
	if err != nil {
		//（ユーザー情報がDBにないとき。ポインタだからnilも許容）
		return nil, errors.New("Nonexistent User")
	}

	fmt.Println(User)
	return &User, nil
}
