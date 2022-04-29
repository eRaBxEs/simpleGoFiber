package main

import (
	"fiber-stats/util"
	"fmt"
)

func main() {

	data := util.Data{}
	_, err := data.LoadConfiguration("util/config.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	dataString := data.GetString("hosts")
	if err != nil {
		fmt.Println(err.Error())
	}
	dataInt := data.GetInt("port")

	dataString2 := data.GetString("database.user")

	dataString3 := data.GetInt("database.port")

	fmt.Println("dataString:", dataString)
	fmt.Println("dataInt:", dataInt)
	fmt.Println("dataString2:", dataString2)
	fmt.Println("dataString3:", dataString3)
}
