package ApiService

import (
	"io"
	"net/http"
)

func Get(url string, params map[string]string, headers map[string]string) (response string, e error) {
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
		return "", err;
	}

	resBody, _ := io.ReadAll(res.Body);
	resBodyString := string(resBody);

	return resBodyString, nil;
}

// func Post() string {

// }