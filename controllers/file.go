package controllers

import (
	"assignment3/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func RegenFile() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	data := structs.FileJson{
		Status: structs.Item{
			Water: rand.Intn(max-min+1) + min,
			Wind:  rand.Intn(max-min+1) + min,
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func ReadFile() structs.FileJson {
	var file structs.FileJson
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &file)
	return file
}
