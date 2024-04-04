package orm

import "os"

// redeclare writer
// func NewWriter(w io.Writer) (writer *csv.Writer) {
//     writer = csv.NewWriter(w)
//     writer.Comma = '\t'

//     return
// }

func SavePage(filename string, html string) error {

	path := "files/pages/" + filename

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(html)
	if err != nil {
		return err
	}

	return nil

}