package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//flag path
	path := flag.String("p", "", "File path to grep")

	//flag search
	search := flag.String("s", "", "Search string from the path file")

	//Define flag!
	flag.Parse()
	files, err := ioutil.ReadDir(*path) //pointer to use it on heap
	if err != nil {
		log.Fatal(err)
	}
	//for each files in directory
	for _, file := range files {
		if file.Mode().IsDir() {
			continue
		}
		// Read file inside directory
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		//details mime of files
		filext := http.DetectContentType(data)
		short := strings.Split(string(data), "\n")
		linum := 1 //line start from 1
		for _, line := range short {
			// indexing substring using the cli
			if strings.Index(line, *search) > -1 {
				if strings.Index(filext, "text/plain") > -1 {
					fmt.Printf("[%s] Line %d:\t %s\n", file.Name(), linum, line)
				} else {
					fmt.Println("Nil")
					break
				}
			}
			linum++
		}
	}
}
