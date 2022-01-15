package core

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

const baseUrl = "https://yts.mx/api/v2/list_movies.json"

// Code from https://mailazy.com/blog/http-request-golang-with-best-practices/ tutorial

func SearchMovie(movie string) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("query_term", movie)

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	fmt.Println(q.Encode())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}