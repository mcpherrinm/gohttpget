package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "expected 1 argument: url to GET")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching: %v", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
