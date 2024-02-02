package chapter1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetUrl() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if !strings.HasPrefix(url, prefix) {
			fmt.Println("Adding prefix.")
			url = prefix + url
		}
		rsp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(rsp.Status)
		_, err = io.Copy(os.Stdout, rsp.Body)
		rsp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
