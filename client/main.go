package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	authValue := "autenticacao"
	// Get request
	getEvenResponse, err := http.Get(fmt.Sprintf("http://localhost:8080/pares&auth=%s", authValue))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	baseURL := "http://localhost:8080/impares"
	params := url.Values{}
	params.Add("auth", authValue)

	u, _ := url.ParseRequestURI(baseURL)
	u.RawQuery = params.Encode()
	urlString := fmt.Sprintf("%v", u)
	// Post request
	getOddResponse, err := http.Get(urlString)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	evenValue, err := unmarshal(getEvenResponse)
	if err != nil {
		fmt.Printf("error parsing response: %s\n", err)
	}

	oddValue, err := unmarshal(getOddResponse)
	if err != nil {
		fmt.Printf("error parsing response: %s\n", err)
	}

	fmt.Println(evenValue, oddValue)
}

func unmarshal(res *http.Response) (parsedInt int, err error) {
	responseByte, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(responseByte, &parsedInt)
	if err != nil {
		return 0, err
	}

	return parsedInt, nil
}
