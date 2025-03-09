package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func scanner() {
	resp, err := http.Get("https://www.hudarashid.com/")
	if err != nil {
		fmt.Printf("There is an error here %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("There is an error here  to scan the here %v\n", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, Let's GO! 😎\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	//scanner()
	fmt.Print("🚙 Starting server on port 8090...\n")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", headers)

	http.ListenAndServe(":8090", nil)

}
