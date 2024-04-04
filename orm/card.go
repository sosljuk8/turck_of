package orm

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCsv(path string, product []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	defer writer.Flush()

	writer.Write(product)

	return nil
}