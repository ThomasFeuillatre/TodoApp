package main

import (
	"TodoApp/Config"
	"TodoApp/Models"
	"TodoApp/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", "mysql://pc86waepw158hqgw:s8uoz9qu4k5elrhj@ryfqldzbliwmq6g5.chr7pe7iynqr.eu-west-1.rds.amazonaws.com:3306/mjurrpat6gcq8pt1")

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()

	r.Run()
}
