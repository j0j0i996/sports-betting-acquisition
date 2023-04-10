package acquisition

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func buildRequest(endpoint string) *http.Request {
	url := "https://v3.football.api-sports.io/" + endpoint
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("x-apisports-key", os.Getenv("XAPISPORTSKEY"))
	return req
}

func executeRequest(req *http.Request) []byte {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

// GetData with optional API Params
func GetData(endpoint string, parameter_map ...map[string]string) []byte {
	req := buildRequest(endpoint)

	if len(parameter_map) != 0 {
		q := req.URL.Query()
		for _, item := range parameter_map {
			for key, value := range item {
				q.Add(key, value)
			}
		}
		req.URL.RawQuery = q.Encode()
	}

	return executeRequest(req)
}
