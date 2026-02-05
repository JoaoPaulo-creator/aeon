package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./teste.aeon")
	if err != nil {
		panic(err.Error())
	}

	formattedData := strings.Fields(string(data))

	for i := range formattedData {
		fmt.Println(formattedData[i])
	}
}
