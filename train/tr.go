package train

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseSample() {
	fileContent, err := ioutil.ReadFile("train/tr.html")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string   https://www.turck.de/
	//text := string(fileContent)
	//fmt.Println(text)

	// Use strings.NewReader to create a reader from the string
	//reader := strings.NewReader(text)

	// Use goquery.NewDocumentFromReader to read the reader
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(fileContent)))
	if err != nil {
		log.Fatal(err)
	}

	// Find category from div.breadcrumbs ul (li a text +|+ li a text+|+li a text)
	category := ""
	doc.Find("div.breadcrumb ul li a").Each(func(i int, s *goquery.Selection) {

		category += "|" + strings.TrimSpace(s.Text())
	})
	category = strings.Replace(category, "|Products|", "", -1)

	// find img from div#compare div.dtimg img (attribute "src" value)
	img := doc.Find("div#compare div.dtimg img").AttrOr("src", "")
	img = "https://www.turck.de" + img

	description := ""

	props := make(map[string]string)
	// find properties from table.tableProd tbody tr; name td first  value td second
	doc.Find("div#infotable1 table.tableProd tbody tr").Each(func(i int, s *goquery.Selection) {
		key := strings.TrimSpace(s.Find("td").First().Text())
		value := strings.TrimSpace(s.Find("td").Eq(1).Text())
		props[key] = value
	})

	properties := ""

	// if properties is not empty then convert map[string]string
	if len(props) > 0 {
		// properties = props to json
		propertiesJSON, err := json.Marshal(props)
		if err != nil {
			log.Println(err)
		} else {
			properties = string(propertiesJSON)
		}
	}

	fmt.Println(category)
	fmt.Println(img)
	fmt.Println(properties)
	fmt.Println(description)
}
