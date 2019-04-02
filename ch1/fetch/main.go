package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		prefix:="http://"
		if !strings.HasPrefix(url, prefix){
			url = prefix + url
		}

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v/n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch read %s: %v/n", url, err)
		//	os.Exit(1)
		//}

		if written, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch read %s: %v/n", url, err)
			os.Exit(1)
		}else{
			fmt.Printf("Writen: %d bytes, HTTP Status: %s", written, resp.Status)
		}
	}
}
