package main

import (
	"fmt"
	"log"
	"net/http"

	"golinkshortener/src/provider"
)

type ServerHandler struct {
	UrlProvider provider.UrlProvider
}

func (handler ServerHandler) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	url, _ := handler.UrlProvider.Get("")
	redirect(resWriter, url)
}

func redirect(resWriter http.ResponseWriter, url string) {
	resWriter.Header().Add("Location", url)
	resWriter.WriteHeader(http.StatusPermanentRedirect)
}

func main() {
	fmt.Println("Server Started")

	s := &http.Server{
		Addr: ":9090",
		Handler: ServerHandler{
			UrlProvider: provider.YamlUrlProvider{
				FilePath: "./link.yaml",
			},
		},
	}

	log.Fatal(s.ListenAndServe())
}
