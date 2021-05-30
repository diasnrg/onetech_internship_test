package acmp_concurrent

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
//wait group, to wait for all goroutines to execute
	var wg sync.WaitGroup
//set counter to the size of an array (will descement after each go routine)
	wg.Add(len(urls))

	values := make(map[string]float64)
	for _,url := range urls {
//running the function for each single url in separate goroutine
		go Difficulty(url, values, &wg)
	}

//waiting for all goroutines and return the map
	wg.Wait()
	return values
}

func Difficulty(url string, values map[string]float64, wg *sync.WaitGroup) {
//decrement the value of wait group in the end
	defer wg.Done()

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
		values[url] = -1
		return
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
		values[url] = -1
		return
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

//syns.Mutex given the access to one goroutine at a time (to reduce 'concurrent map write error')
	var mutex = &sync.Mutex{}
	mutex.Lock()

	values[url] = d

	mutex.Unlock()
}