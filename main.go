package main

import (
	"flag"
	"fmt"
	"sezzle/data"
)

func main() {
	var input = flag.String("in", "", "How would you like to save this data?")
	flag.Parse()

	db := data.NewDatabase()
	service := data.NewDataService(db, data.DataTypeEnum(*input))
	result, err := service.Save()
	if err != nil {
		fmt.Println(fmt.Errorf("save failed: %v", err))
		return
	}
	fmt.Println("Result: ", result)
}
