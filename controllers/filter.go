package controllers

import (
	"assignment3/structs"
	"fmt"
)

func Filter(json structs.Item) structs.Response {
	var statusWind string
	var statusWater string
	var response structs.Response
	if json.Water <= 5 {
		statusWater = "aman"
	} else if json.Water >= 6 && json.Water <= 8 {
		statusWater = "siaga"
	} else if json.Water > 8 {
		statusWater = "bahaya"
	}

	if json.Wind <= 6 {
		statusWind = "aman"
	} else if json.Wind >= 7 && json.Wind <= 15 {
		statusWind = "siaga"
	} else if json.Wind > 15 {
		statusWind = "bahaya"
	}

	response.Response = "Water : " + fmt.Sprint(json.Water) + " meter => " + statusWater + " | Wind : " + fmt.Sprint(json.Wind) + " meter/detik => " + statusWind
	return response
}
