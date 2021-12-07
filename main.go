package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// makeCall will make an HTTP GET or POST to the specified url and print the response body.
func makeCall(v, url, b string) {

	readBody := func(resp *http.Response) []byte {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		return b
	}

	if v == "POST" {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(b)))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(readBody(resp)))
	} else {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(readBody(resp)))
	}
}

// handleInputs parses commandline flags for usage
func handleInputs() (int, string, string, string) {
	c := flag.Int("c", 1, "The `number of requests` to make per second.")
	x := flag.String("x", "GET", "The `HTTP verb` to use. [GET|POST].")
	url := flag.String("u", "https://google.com", "The `URL to hit`.")
	b := flag.String("b", "{}", "The `JSON body` to send with a -x POST.")

	flag.Parse()

	// if -x POST then body must not be empty and exits
	if *x == "POST" {
		if *b == "{}" {
			log.Fatalln("no body specified")
		}
	}

	// print flag usage if not enough inputs and exits
	if flag.NFlag() < 2 {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
		log.Fatalln("not enough inputs")
	}

	return *c, *x, *url, *b
}

func main() {
	count, method, url, body := handleInputs()

	// continual loop to make c requests per one second
	for {
		for i := 0; i < count; i++ {
			go makeCall(method, url, body)
		}
		time.Sleep(time.Second)
	}
}
