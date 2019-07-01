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
			"ID":            "goya",
			"Password":      "syogo",
			"Name":          "DESTINY-ZERO8",
			"Email":         "goya@goya.com",
			"StreetAddress": "沖縄県呉屋市呉屋町5858-5",
			"IsAdmin":       "false",
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
		password, _ := info["Password"]
		name, _ := info["Name"]
		email, _ := info["Email"]
		streetAddress, _ := info["StreetAddress"]
		isAdmin, _ := strconv.ParseBool(info["IsAdmin"])
		createUser(model.User{
			ID:            string(ID),
			Password:      string(password),
			Name:          string(name),
			Email:         string(email),
			StreetAddress: string(streetAddress),
			IsAdmin:       bool(isAdmin),
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
