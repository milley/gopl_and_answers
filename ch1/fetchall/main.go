// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open file error: %v\n", err)
		os.Exit(1)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		//	fmt.Println(<-ch) // receive from channel ch
		//fmt.Fprintf(file, "%s\n", <-ch)
		if _, err := file.Write([]byte(fmt.Sprintf("%s\n", <-ch))); err != nil {
			fmt.Printf("Write url response time to file error: %v\n", err)
		}
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
	if _, err := file.Write([]byte(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))); err != nil {
		fmt.Printf("Write total response time to file error: %v\n", err)
	}

	if err := file.Close(); err != nil {
		fmt.Printf("Close file error: %v\n", err)
		os.Exit(1)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
