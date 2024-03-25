package ApiService

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Get(url string, params map[string]string, headers map[string]string) (responseBytes []byte, e error) {
	client := &http.Client{};
	req, _ := http.NewRequest(http.MethodGet, url, nil);
	queryParams := req.URL.Query()

	for key, value := range params{
		queryParams.Add(key, value)
	}
	req.URL.RawQuery = queryParams.Encode();

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req);

	if err != nil {
		return nil, err;
	}

	resBody, _ := io.ReadAll(res.Body);

	return resBody, nil;
}

func Post(url string, body interface{}, headers map[string]string) (responseBytes []byte, e error) {
	client := &http.Client{};
	bodyStr := fmt.Sprintf("%v", body);
	bytesBody := []byte(bodyStr)
	// jsonValue, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bytesBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	// for key, value := range headers {
	// 	req.Header.Set(key, value)
	// }

	res, err := client.Do(req);
	if err != nil {
		log.Fatal(err)
	}

	resBytes, _ := io.ReadAll(res.Body);

	return resBytes, nil;
}