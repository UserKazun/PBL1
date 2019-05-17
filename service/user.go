package service

import (
	"errors"
	"fmt"

	"github.com/PBL1/model"
)

// CreateUser ...DBに与えられたデータをinsertする
func CreateUser(user model.User) (model.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CheckAccount ...ログインされたアカウント情報が存在するか確かめる
func CheckAccount(userID string, password string) (*model.User, error) {
	User := model.User{}

	err := db.Where("id = ? and password = ?", userID, password).First(&User).Error
	if err != nil {
		//（ユーザー情報がDBにないとき。ポインタだからnilも許容）
		return nil, errors.New("Nonexistent User")
	}

	fmt.Println(User)
	return &User, nil
}

func PostTrueToAdmin(userID string) {
	userBefore := model.User{}

	userBefore.ID = userID
	userAfter := userBefore

	db.First(&userAfter)

	userAfter.IsAdmin = true

	// 更新を実行
	db.Model(&userBefore).Update(&userAfter)

}
