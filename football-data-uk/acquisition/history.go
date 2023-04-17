package acquisition

import (
	"log"
	"reflect"
	"strconv"
	"time"
)

type FixtureItem struct {
	Time     time.Time
	Div      string
	HomeTeam string
	AwayTeam string
	B365H    float32
	B365A    float32
	B365D    float32
}

// FetchData requests csv from URL and saves / returns it as string
// example url: https://www.football-data.co.uk/mmz4281/1718/N1.csv
func fetchHistoricData(year_code string, league string) ([][]string, error) {

	url := BASE_URL + "/" + HIST_ODD_ENDPOINT + "/" + year_code + "/" + LeagueMap[league] + ".csv"
	data, err := ReadCSVFromUrl(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// stringInSlice checks if a single string is in a list of strings
func isStringInSlice(str_to_check string, slice []string) bool {
	for _, b := range slice {

		if b == str_to_check {
			return true
		}
	}
	return false
}

// filterColumnsByName filters an array of array by only keeping the columns
// in which the name in the first row matches one of the col_names_to_keep input
// It retuns the data as map, in which the key is the column name
func modelData(data [][]string) []FixtureItem {

	// Get index of columns
	idx_map := make(map[string]int)

	// Get column names to keep
	cols_to_keep := getFieldNamesFromFixtureItem(FixtureItem{})

	// Find column number of cols_to_keep in returned csv
	for index, element := range data[0] {

		if isStringInSlice(element, cols_to_keep) {
			idx_map[element] = index
		}
	}
	if len(cols_to_keep) != len(idx_map) {
		log.Fatalln("One of the excpected column names is not part in csv header")
	}

	// Model data
	var modelled_data []FixtureItem
	for i := 1; i < len(data); i++ {
		fixture := FixtureItem{}
		for fieldName, index := range idx_map {
			v := reflect.ValueOf(&fixture).Elem()
			f := v.FieldByName(fieldName)

			if f.IsValid() {
				if f.CanSet() {
					if f.Type().String() == "float32" {
						float_number, err := strconv.ParseFloat(data[i][index], 32)
						if err != nil {
							log.Fatalln("Could not convert String to Float")
						}
						f.SetFloat(float_number)
					} else if f.Type().String() == "string" {
						f.SetString(data[i][index])
					} else {
						log.Println("Unknown Field type")
					}
				} else {
					log.Fatalln("Field is not settable")
				}
			} else {
				log.Fatalln("Field not found")
			}
		}
		modelled_data = append(modelled_data, fixture)
	}

	return modelled_data
}

func GetHistoricData(year_code string, league string) []FixtureItem {
	// Get Data from URL as slcie of slices, with first array being the header
	data, err := fetchHistoricData(year_code, league)
	if err != nil {
		log.Fatal(err)
	}

	modelledData := modelData(data)

	return modelledData
}
