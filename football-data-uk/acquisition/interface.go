package acquisition

import (
	"encoding/csv"
	"log"
	"net/http"
)

// FetchData requests csv from URL and saves / returns it as string
func ReadCSVFromUrl(url string) ([][]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
