package main

import (
	//"fmt"
	"bytes"
	//"errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"launchpad.net/xmlpath"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	var url, body string

	url = getattribute(os.Args)

	if len(strings.Trim(url, " ")) < 1 {
		log.Println("Empty parameter provided")
		os.Exit(0)
	}

	log.Println("Starting query...")

	//"http://yandex.ru/yandsearch?lr=10747&text=11&suggest_reqid=842028564142219026603767268998476"

	//reader := strings.NewReader(string(body))
	//log.Println(string(body))

	body = Query(url)
	dumpHmtl([]byte(body))
	ParsePage(body)
	ParsePagin(body)
}

func ParsePage(text string) {
	reader := strings.NewReader(text)
	root, err := html.Parse(reader)

	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	html.Render(&b, root)
	fixedHtml := b.String()

	reader = strings.NewReader(fixedHtml)
	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	regPath := xmlpath.MustCompile("/html/body/div[2]/div[3]/div/div[4]/div[2]/div")
	if value, ok := regPath.String(xmlroot); ok {
		log.Println(value)
	}

	path := xmlpath.MustCompile("/html/body/div[2]/div[3]/div/div[3]/div[1]/div[1]/div[2]/div/div/div/h2/a/@href")

	iter := path.Iter(xmlroot)
	for iter.Next() {
		url := iter.Node().String()
		log.Println(url)
	}

	//paginPath := xmlpath.MustCompile("/html/body/div[2]/div[3]/div/div[4]/div[1]/div/span[1]")
}

func ParsePagin(text string) {
	reader := strings.NewReader(text)
	root, err := html.Parse(reader)

	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	html.Render(&b, root)
	fixedHtml := b.String()

	reader = strings.NewReader(fixedHtml)
	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	paginPath := xmlpath.MustCompile("/html/body/div[2]/div[3]/div/div[4]/div[1]/div/span/a/@href")
	if value, ok := paginPath.String(xmlroot); ok {
		log.Println(value)
	}
}

func Query(url string) string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:26.0) Gecko/20100101 Firefox/26.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func getattribute(params []string) string {
	if len(os.Args) < 2 {
		log.Println("Missing url parameter")
		os.Exit(0)
	} else {
		return os.Args[1]
	}

	return ""
}

func dumpHmtl(text []byte) {

	ioutil.WriteFile("output.html", text, 0644)

}
