package core

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"

	// "strconv"
	// "github.com/manifoldco/promptui"
	// "errors"

)

const baseUrl = "https://yts.mx/api/v2/list_movies.json"

type Movie struct {
	Id int `json:id`
	Url string `json:url`
	ImdbCode string `json:imdb_code`
	Title string `json:title`
	Year int `json:year`
	Summary string `json:summary`
}


type MovieData struct {
	Count int `json:movie_count`
	Limit int `json:limit`
	Page int `json:page_number`
	Movies []Movie `json:movies`
}

type YTSResponse struct  {
	Status string `json:string`
	Message string `json:status_message`
	Data MovieData `json:data`
}


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
	if(resp.Status != "200 OK") {
		log.Print("Error reaching api")
		return
	}
	var yts YTSResponse
	json.Unmarshal(responseBody, &yts)
	//fmt.Println(string(responseBody))

	for _, movie:= range yts.Data.Movies {
		fmt.Printf("Title:%s, Year:%d , Summary:%s \n",movie.Title, movie.Year, movie.Summary)
	}


	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)
	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	// prompt := promptui.Prompt{
	// 	Label:    "Number",
	// 	Validate: validate,
	// }

	// result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// fmt.Printf("You choose %q\n", result)

}