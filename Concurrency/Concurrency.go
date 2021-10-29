package main

import (
	"log"
	"net/http"
)

type Lulusan struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}

type Result struct {
	Resource_id string    `json:"resource_id"`
	Fields      []Fields  `json:"fields"`
	Records     []Records `json:"records"`
}

type Fields struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Records struct {
	_Id           int    `json:"_id"`
	Sex           string `json:"sex"`
	JumlahLulusan string `json:"no_of_graduates"`
	Jurusan       string `json:"type_of_course"`
	Tahun         string `json:"year"`
}

func main() {
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
}
