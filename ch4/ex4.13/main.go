package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// URL omdb api url
const URL = "http://www.omdbapi.com/?t="

// Movie struct
type Movie struct {
	Response string
	Error    string
	Poster   string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: need to input keywords")
		os.Exit(1)
	}

	keywords := url.QueryEscape(strings.Join(os.Args[1:], " "))
	resp, err := http.Get(URL + keywords)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	if movie.Response != "True" {
		fmt.Println(movie.Error)
		os.Exit(4)
	}

	if movie.Poster == "" {
		fmt.Println("No poster")
		os.Exit(5)
	}

	fmt.Println("Downloading file...")
	rawURL := movie.Poster
	fileURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	filename := segments[len(segments)-1]
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp2, err := check.Get(rawURL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp2.Body.Close()
	fmt.Println(resp2.Status)

	size, err := io.Copy(file, resp2.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s with %v bytes downloaded\n", filename, size)
}
