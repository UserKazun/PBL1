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

// GetPasswordByID ...ID別にパスワード情報を取得
func GetPasswordByID(userID string) (string, error) {
	user := model.User{}

	err := db.Where("id = ?", userID).First(&user).Error
	if &user != nil && err != nil {
		return "", err
	}

	return user.Password, nil
}

//PutUserByID ...ユーザ情報の更新
func PutUserByID(userID string, userName string, userEmail string, UserStreetAddress string, newPassword string) error {
	beforeUser := model.User{}

	err := db.Where("id = ?", userID).First(&beforeUser).Error
	if &beforeUser != nil && err != nil {
		return err
	}

	afterUser := beforeUser
	afterUser.ID = userID
	afterUser.Name = userName
	afterUser.Email = userEmail
	afterUser.StreetAddress = UserStreetAddress
	afterUser.Password = newPassword

	err = db.Model(&beforeUser).Update(afterUser).Error
	if err != nil {
		return err
	}

	return nil
}
