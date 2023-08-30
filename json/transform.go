package json

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string `json:"titleee"`
	Year   int    `json:"yearrr"`
	Length string `json:"length,default=222"`
}

func marshal() {
	king := Movie{
		Title: "huanglei",
		Year:  2001,
	}
	data, _ := json.Marshal(king)
	fmt.Println(string(data))
}
