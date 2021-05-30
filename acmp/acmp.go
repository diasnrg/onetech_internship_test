package acmp

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

func Difficulty(url string) float64 {
//create the request object with method - 'get' and passed 'url'
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

//set the cookie for the context in english (by default - russian)
	req.Header.Set("Cookie", "English=1")

//create the 'client' object to make response with request object
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

//if page doesn't exist - return '-1'
	if res.StatusCode != 200 {
		return -1
	}
	
//load HTML document from response's body
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

//find text of the in format (time limit: x sec, memory limit: x MB, difficulty: x%)
	text := doc.Find("td center i").Text()

//if the page is not in the correct format 
	if len(text) == 0 {
		return -1
	}

//take the part in format 'Difficulty: x%'
	difficultyText := regexp.MustCompile(`Difficulty: (\d+)%`).Find([]byte(text))

//traversing the 'Difficulty: x%' and appending to result var in string format if the byte is numberic
	resultStr := ""
	for _,v := range difficultyText {
		if v >= '0' && v <= '9' {
			resultStr += string(v)
		}
	}

//parsing the string into the float64 format
	d, err := strconv.ParseFloat(string(resultStr), 64)
	if err != nil {
		log.Fatal(err)
	}

	return d
}
