package main

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/xmlpath"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("http://yandex.ru/yandsearch?text=%D1%87%D0%B5%D1%82%D0%BA%D0%B8%D0%B5%20%D0%BC%D0%B0%D0%BA%D0%B0%D1%80%D0%BE%D0%BD%D1%8B")
	if err != nil {
		//fmt.Printf(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	reader := strings.NewReader(string(body))

	xmlroot, xmlerr := xmlpath.ParseHTML(reader)

	if xmlerr != nil {
		log.Fatal(xmlerr)
	}

	path := xmlpath.MustCompile("//div[@id=someid]")
	if value, ok := path.String(xmlroot); ok {
		log.Println("Found:", value)
	}

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf(string(body))
	fmt.Printf("hello from golang")

}
