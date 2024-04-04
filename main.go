package main

import (
	"encoding/csv"
	"strings"

	"io"
	"log"
	"os"

	"github.com/sosljuk8/turckof/dto"
	"github.com/sosljuk8/turckof/orm"
	"github.com/sosljuk8/turckof/parse"
)

// func main() {
// 	train.ParseSample()
// }

var CsvPath = "files/parsed/turck.csv"

func init() {
	file, err := os.Create(CsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func main() {
	file, err := os.Open("files/turck_data.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	cards := map[string]*dto.PCard{}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(record[4], "unknown") {
			continue
		}
		card := &dto.PCard{
			Brand:    record[0],
			Model:    record[2],
			Name:     record[3],
			SKU:      record[4],
			Price:    record[5],
			Currency: record[6],
		}

		cards[card.SKU] = card
	}

	for _, card := range cards {
		addata := parse.Parse(card.SKU)
		card.Category = addata["category"]
		card.Img = addata["img"]
		card.Properties = addata["properties"]
		card.Description = addata["description"]
		card.Source = addata["source"]
		card.File = addata["file"]

		err := orm.WriteCsv(CsvPath, card.String())
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(card)
	}
}
