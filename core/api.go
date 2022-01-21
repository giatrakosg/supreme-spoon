package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseUrl = "https://yts.mx/api/v2/list_movies.json"

type Torrent struct {
	Url      string `json:"url"`
	Hash     string `json:"hash"`
	Quality  string `json:"quality"`
	Type     string `json:"type"`
	Seeds    int    `json:"seeds"`
	Peers    int    `json:"peers"`
	Size     string `json:"size"`
	Uploaded string `json:"date_uploaded"`
}

type Movie struct {
	Id       int       `json:"id"`
	Url      string    `json:"url"`
	ImdbCode string    `json:"imdb_code"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Summary  string    `json:"summary"`
	Slug     string    `json:"slug"`
	Torrents []Torrent `json:"torrents"`
}

type MovieData struct {
	Count  int     `json:"movie_count"`
	Limit  int     `json:"limit"`
	Page   int     `json:"page_number"`
	Movies []Movie `json:"movies"`
}

type YTSResponse struct {
	Status  string    `json:"string"`
	Message string    `json:"status_message"`
	Data    MovieData `json:"data"`
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

	if resp.Status != "200 OK" {
		log.Print("Error reaching api")
		return
	}
	var yts YTSResponse
	json.Unmarshal(responseBody, &yts)

	// Show the select movie ui
	SelectMovie(yts.Data.Movies)

}

/*
	We download the torrent file in the provided url and store in ./data
*/

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
