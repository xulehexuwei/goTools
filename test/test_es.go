package main

import (
	"fmt"
	"goTools/db"
)

func main() {
	searchResult := db.EsQueryByMatch("es", "aminer_research", "name_zh", "计量学")
	fmt.Println(searchResult)
}
