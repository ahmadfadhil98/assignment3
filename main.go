package main

import (
	"assignment3/controllers"
	"assignment3/structs"
	"fmt"
	"time"
)

func main() {
	var response structs.Response
	var file structs.FileJson
	for {
		controllers.RegenFile()
		file = controllers.ReadFile()
		response = controllers.Filter(file.Status)
		fmt.Println(response.Response)
		time.Sleep(15 * time.Second)
	}
}
