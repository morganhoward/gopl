// Fetch prints the content found at each specified url

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		u := handleURL(url)
		resp, err := http.Get(u)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", u, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func handleURL(url string) (u string) {
	if !strings.HasPrefix(url, "http://") {
		u = strings.Join([]string{"http://", url}, "")
	} else {
		u = url
	}
	return
}
