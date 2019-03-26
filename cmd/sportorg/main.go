package main

import (
	"fmt"
	"github.com/sportorg/online/api"
	"github.com/sportorg/online/database"
)

func main() {
	db := database.Connect("root:P@ssw0rd@tcp(127.0.0.1:3306)/sportorg")
	defer db.Close()
	fmt.Println("Running Sportorg Online")
	api.Run(":8080", "./web/dist/sportorg-online")
}
