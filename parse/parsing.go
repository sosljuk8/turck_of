package parse

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sosljuk8/turckof/orm"
)

func Parse(sku string) map[string]string {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.turck.de"),
	)

	datar := map[string]string{}

	// After making a request print "Visited ..."
	c.OnResponse(func(r *colly.Response) {

		if bytes.Contains(r.Body, []byte(`<h1>404 Error, Page Not Found</h1>`)) {
			// just skip this url, no errors triggered
			return
		}

		fmt.Println("THIS IS THE PAGE!!!!", r.Request.URL)

		// Parse the page
		data, err := ParsePage(bytes.NewBuffer(r.Body))
		if err != nil {
			log.Println(err)
			return
		}
		datar = data
		datar["source"] = r.Request.URL.String()
		datar["file"] = Hash(sku) + ".html"

		err = orm.SavePage(datar["file"], string(r.Body))
		if err != nil {
			log.Println(err)
			return
		}

		// card.Category = (*data)["category"]
		// card.Img = (*data)["img"]
		// card.Description = (*data)["description"]
		// card.Properties = (*data)["properties"]

	})

	c.Visit("https://www.turck.de/en/product/" + sku)
	return datar
}

func ParsePage(body *bytes.Buffer) (map[string]string, error) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)

	// Find category from div.breadcrumbs ul (li a text +|+ li a text+|+li a text)
	category := ""
	doc.Find("div.breadcrumb ul li a").Each(func(i int, s *goquery.Selection) {

		category += "|" + strings.TrimSpace(s.Text())
	})
	data["category"] = strings.Replace(category, "|Products|", "", -1)

	// find img from div#compare div.dtimg img (attribute "src" value)
	img := doc.Find("div#compare div.dtimg img").AttrOr("src", "")
	data["img"] = "https://www.turck.de" + img

	data["description"] = ""

	props := make(map[string]string)
	// find properties from table.tableProd tbody tr; name td first  value td second
	doc.Find("div#infotable1 table.tableProd tbody tr").Each(func(i int, s *goquery.Selection) {
		key := strings.TrimSpace(s.Find("td").First().Text())
		value := strings.TrimSpace(s.Find("td").Eq(1).Text())
		props[key] = value
	})

	data["properties"] = ""

	// if properties is not empty then convert map[string]string
	if len(props) > 0 {
		// properties = props to json
		propertiesJSON, err := json.Marshal(props)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		data["properties"] = string(propertiesJSON)
	}

	//fmt.Println(data)

	return data, nil
}

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
