package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)
   
   func main() {
	file, err := os.Open("files/turck.csv")
   
	if err != nil {
		   log.Fatal(err)
	   }
	defer file.Close()
   
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1


	//cards := map[string]*dto.PCard{}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[4])
	}
   }
