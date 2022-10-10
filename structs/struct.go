package structs

type Item struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type FileJson struct {
	Status Item `json:"status"`
}

type Response struct {
	Response string `json:"response"`
}
