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

// PostTrueToIsAdmin ...対象のuserTableのIsAdminをTrueに変更
func PostTrueToIsAdmin(userID string) {
	userBefore := model.User{}

	userBefore.ID = userID
	userAfter := userBefore

	db.First(&userAfter)

	userAfter.IsAdmin = true

	// 更新を実行
	db.Model(&userBefore).Update(&userAfter)
}

// PostFalseToIsAdmin ...対象のuserTableのIsAdminをFalseに変更
func PostFalseToIsAdmin(userID string) {
	userBefore := model.User{}

	userBefore.ID = userID
	userAfter := userBefore

	db.First(&userAfter)

	userAfter.IsAdmin = false

	// 更新を実行
	// db.Model(&userBefore).Update(&userAfter)
	db.Save(&userAfter)
}

// GetUserByID ...ID別にユーザ情報を取得
func GetUserByID(userID string) (*model.User, error) {
	user := model.User{}

	err := db.Where("id = ?", userID).First(&user).Error
	if &user != nil && err != nil {
		return nil, err
	}

	return &user, nil
}
