package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golinkshortener/src/provider"
)

type ServerHandler struct {
	UrlProvider provider.UrlProvider
}

func (handler ServerHandler) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	trimmedPath := req.URL.Path[1:]

	fmt.Printf("Query for %s\n", trimmedPath)
	url, err := handler.UrlProvider.Get(trimmedPath)

	if err == nil {
		fmt.Printf("Redirecting for %s to %s\n", trimmedPath, url)
		redirect(resWriter, url)
	} else {
		fmt.Printf("Error query for %s\n", trimmedPath)
	}
}

func redirect(resWriter http.ResponseWriter, url string) {
	resWriter.Header().Add("Location", url)
	resWriter.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	fmt.Println("Server Started")

	linkPath := flag.String("path", "./link.yaml", "a file path to link yaml")
	flag.Parse()

	fmt.Printf("Loading YAML path from: %s", *linkPath)

	s := &http.Server{
		Addr: ":9090",
		Handler: ServerHandler{
			UrlProvider: provider.YamlUrlProvider{
				FilePath: *linkPath,
			},
		},
	}

	log.Fatal(s.ListenAndServe())
}
