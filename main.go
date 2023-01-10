package main

import (
	"fmt"
	"goadmin/dao/mysql"
	"goadmin/routers"
	"goadmin/utils"
)

func main() {

	g := routers.Router()
	err := mysql.Initdb()
	if err != nil {
		fmt.Println("Db连接错误", err)
		return
	}

	if err := utils.Init("2020-07-01", int64(1)); err != nil {
		fmt.Println("Init snowflake err :", err)
		return
	}

	//go utils.SCronSta()

	g.Run("127.0.0.1:8001")
}
