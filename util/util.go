package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGetRequest(client *http.Client, url string, args ...string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []byte{}, errors.New("could not use the URL")
	}

	q := req.URL.Query()

	for i := 0; i < len(args)-1; i += 2 {
		q.Add(args[i], args[i+1])
	}

	req.URL.RawQuery = q.Encode()

	resp, err := (*client).Do(req)

	if err != nil {
		return []byte{}, errors.New("could not get response")
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func PostJSON(client *http.Client, url string, payload interface{}) ([]byte, error) {

	payloadJSON, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))

	if err != nil {
		fmt.Println("Error creating new request instance:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured while sending request:", err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
