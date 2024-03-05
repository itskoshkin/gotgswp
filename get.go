package gotgswp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {

		return "", fmt.Errorf("http.Get returned an error: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {

		return "", fmt.Errorf("io.ReadAll returned an error: %v", err)
	}

	return string(body), nil
}
