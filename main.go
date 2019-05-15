package main

import (
	"github.com/PBL1/router"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
