package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
)

// CreateSeedUsers ...ここで記入したデータをDBにinsertする
func CreateSeedUsers() {
	usersInfos := []map[string]string{
		map[string]string{
			"ID":       "goya",
			"Password": "syogo",
			"Name":     "destinyZero",
			"IsAdmin":  "false",
		},
		map[string]string{
			"ID":       "Kmyan",
			"Password": "opai1919",
			"Name":     "かまやん",
			"IsAdmin":  "false",
		},
	}

	for _, info := range usersInfos {
		ID, _ := info["ID"]
		Password, _ := info["Password"]
		Name, _ := info["Name"]
		IsAdmin, _ := strconv.ParseBool(info["IsAdmin"])
		createUser(model.User{
			ID:       string(ID),
			Password: string(Password),
			Name:     string(Name),
			IsAdmin:  bool(IsAdmin),
		})
	}
}

// createUser ...渡されたデータをserviceの関数へ渡し、DBにinsertさせる
func createUser(user model.User) {
	user, err := service.CreateUser(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user created\n")
}
